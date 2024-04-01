package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello from greetings"

	fmt.Println(strings.Contains(greeting, "Hello!"))
	fmt.Println(strings.ReplaceAll(greeting,"Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "lo"))
	fmt.Println(strings.Split(greeting, "  "))
	fmt.Println("Original", greeting)
}
