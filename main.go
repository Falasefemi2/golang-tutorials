package main

import (
	"fmt"
)

func main() {
	var ages = [5]int{1,2,3,4,5}
	fmt.Println(ages, len(ages))

	// slices
	var scores = []int{100,90,80}
	fmt.Println(scores)
	scores[2] = 70
	fmt.Println(scores)
	newvalue := append(scores, 60)
	fmt.Println(newvalue)

	// slices ranges
	rangeOne :=   ages[:4]
	fmt.Println(rangeOne)

	rangeOne = append(rangeOne, 7)
	fmt.Println(rangeOne)


}