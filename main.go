package main

import (
	"fmt"
)

func main() {
	// greeting := "Hello my friend"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// // fmt.Println(strings.Index(greeting, "y"))
	// fmt.Println(strings.Split(greeting, " "))
	// fmt.Println("original", greeting)


	// ages := []int{20,302,22,11,30,990,2202}

	// sort.Ints(ages)
	// fmt.Println(ages)

	var ages map[string]int
	ages = make(map[string]int)

	ages["femi"] = 24
	ages["sam"] = 26
	ages["ayo"] = 27

	femiAge := ages["femi"]
	fmt.Println(femiAge)

	if age, ok := ages["sam"]; ok {
		fmt.Println("sam age is", age)
	} else {
		fmt.Println("sam age is not available")
	}

}