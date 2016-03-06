package main.java;

public class HelloWorld {

    public static String hello(String name) {
        if (name == null || name.isEmpty()) {
            return hello();
        } else {
            return "Hello, " + name + "!";
        }
    }

    public static String hello() {
        return "Hello, World!";
    }
}
