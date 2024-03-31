package main

import (
	"fmt"
)

func main() {
	// var ages  [3]int = [3]int{1,2,3}
	var ages = [3]int{1,2,3}

	names := [4]string{"femi", "samuel", "dare", "ayo"}
	names[1] = "yinka"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices
	var scores = []int{100,90,80}
	fmt.Println(scores)
	scores[2] = 34
	scores = append(scores, 67)
	fmt.Println(scores)
}