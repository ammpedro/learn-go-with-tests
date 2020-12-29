package main

import "fmt"

const engHelloPrefix = "Hello, "

// Hello returns Hello, world
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return engHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
