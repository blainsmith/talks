Class Person {
    String firstName;
    String lastName;
    int age;

    public void name() String {
        return firstName + " " + lastName;
    }
}

Class Designer extends Person {
    // Designer-related properties
}

Class Developer extends Person {
    // Designer-related properties
}

Person john = new Designer();
john.name();
Person sally = new Developer();
sally.name();