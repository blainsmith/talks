interface Person {
    void name() String;
}

Class Designer implements Person {
    String name;
    public String name() {
        return name;
    }
}

Class Developer implements Person {
    String firstName;
    String lastName;
    public String name() {
        return firstName + " " + lastName;
    }
}

Person eric = new Designer();
eric.name();
Person sally = new Developer();
sally.name();