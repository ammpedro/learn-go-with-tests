package main

import "fmt"

const es = "es"
const fr = "fr"
const engHelloPrefix = "Hello, "
const esHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

// Hello returns Hello, world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case fr:
		prefix = frHelloPrefix
	case es:
		prefix = esHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
