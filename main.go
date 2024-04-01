package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello from greetings"

	// fmt.Println(strings.Contains(greeting, "Hello!"))
	// fmt.Println(strings.ReplaceAll(greeting,"Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "lo"))
	// fmt.Println(strings.Split(greeting, "  "))
	// fmt.Println("Original", greeting)

	ages := []int{10, 20, 30, 40, 50, 60, 80, 70, 789}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90)
	fmt.Println(index)

	names := []string{"femi","sam","ayo","dare"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "femi"))
}
