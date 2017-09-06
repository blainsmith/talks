// START decl OMIT
package main

import "fmt"

type Person interface {
	Name() string
}

type Designer struct {
	FirstName string
	LastName  string
}

func (d Designer) Name() string {
	return fmt.Sprintf("%s %s", d.FirstName, d.LastName)
}

type Developer int

func (d Developer) Name() string {
	return "Sorry, developers have no names."
}

// END decl OMIT

// START main OMIT
func MyNameIs(p Person) {
	name := p.Name()
	fmt.Println(name)
}

func main() {
	var designer = Designer{FirstName: "Eric", LastName: "Strowb"}
	var developer = Developer(87)

	MyNameIs(designer)
	MyNameIs(developer)
}

// END main OMIT