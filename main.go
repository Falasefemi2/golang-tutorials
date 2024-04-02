package main

import "fmt"

func getTheSmallest() {
	x := []int {
		20,5,4,1,-1,
	}

	if len(x) == 0 {
		fmt.Println("Slice is empty")
		return
	}

	smallest := x[0]
	for _, num := range x {
		fmt.Println(num)
		if num < smallest {
			smallest = num
		}
	}
	fmt.Println("The smallest number is:", smallest)

}

func main() {
	getTheSmallest()
}
