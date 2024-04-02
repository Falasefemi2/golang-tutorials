package main

import (
	"fmt"
)


func getEvenOdd(numbers []int) ([]int, []int) {
	var evenNumber []int
	var oddNumber []int

	for _, v := range numbers {
		if v % 2 == 0 {
			evenNumber = append(evenNumber, v)
		} else {
			evenNumber = append(oddNumber, v)

		}
	}
	return evenNumber, oddNumber
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even, odd := getEvenOdd(numbers)
	fmt.Println("Even numbers:", even)
    fmt.Println("Odd numbers:", odd)
}