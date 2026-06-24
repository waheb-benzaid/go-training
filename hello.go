package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("waheb"))
}

func Hello(name string) string {
	if name == "" {
		return "Hello, world"
	}
	return englishHelloPrefix + name
}
