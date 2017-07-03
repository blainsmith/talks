package main

import (
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"gitlab.com/oddnetworks/feed-generator-service/rgs"
	"gitlab.com/oddnetworks/feed-generator-service/store"
)

// START main OMIT
func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	wps, _ := strconv.Atoi(os.Getenv("WORKER_POOL_SIZE"))
	ps, _ := strconv.Atoi(os.Getenv("POOL_SLEEP"))

	dynamoStore := store.NewDynamoStore()

	errc := make(chan error)

	// Start the Amazon FireTV pool
	go func() {
		errc <- acd.StartPool(wps, ps, logger, dynamoStore)
	}()

	// Start the Roku Global Search pool
	go func() {
		errc <- rgs.StartPool(wps, ps, logger, dynamoStore)
	}()

	logger.Log("exit", <-errc)
}

// END main OMIT

// START StartPool OMIT
func StartPool(workers, sleep int, logger log.Logger, db store.Store) error {
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "pool", "roku_global_search")

	var platforms []store.Platform
	chanPlatforms := make(chan store.Platform)

	// Start a bunch of workers
	for w := 1; w <= workers; w++ {
		go Worker(chanPlatforms, logger, db)
	}

	// Endless loop that sleeps
	for {
		// Scan the db for all roku platforms and send those across the channel to the worker pool
		db.Scan(store.Key{Type: "platform"}, &platforms, "contains($, ?)", "platform", "roku")
		for _, platform := range platforms {
			chanPlatforms <- platform
		}
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

// END StartPool OMIT
// START Worker1 OMIT
func Worker(platforms <-chan store.Platform, logger log.Logger, db store.Store) error {
	// Read platforms off the channel and start generating the feed
	for platform := range platforms {
		feed := &RGSFeed{} // Create the feed struct with XML struct tags

		collectionIds := make([]string, 0) // Seasons, series, playlists, etc
		videoIds := make([]string, 0)      // Movies, shows, shorts, clips, etc.

		for _, view := range platform.Views {
			for key := range view.Relationships {
				relationData := view.Relationships[key]["data"]
				for _, item := range relationData {
					if item.Type == "collection" {
						collectionIds = append(collectionIds, item.ID)
					} else if item.Type == "video" {
						videoIds = append(videoIds, item.ID)
					}
				}
			}
		}
		// END Worker1 OMIT
		// START Worker2 OMIT
		// Create a Pipe to connect the XML encoder with the S3 writer
		feedReader, feedWriter := io.Pipe()

		// Needs to be in a goroutine so it does not deadlock
		go func() {
			// Close the writer when it is done
			defer feedWriter.Close()

			// Encode the feed to the writer end of the Pipe
			feedWriter.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>"))
			xml.NewEncoder(feedWriter).Encode(feed)
		}()

		// Write the writer end of the Pipe to S3
		store.S3.PutObjectStreaming("odd-feeds", "rgs/"+platform.Channel+".xml", feedReader)

	}

	return nil
}

// END Worker2 OMIT
