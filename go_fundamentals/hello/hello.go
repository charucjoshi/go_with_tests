package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
func Hello(name string) string {
	if len(name) == 0 {
		name = "World"	
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Charu"))
}
