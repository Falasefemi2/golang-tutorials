package main

import (
	"fmt"
)

func main() {
	// var ages  [3]int = [3]int{1,2,3}
	// var ages = [3]int{1,2,3}

	// names := [4]string{"femi", "samuel", "dare", "ayo"}
	// names[1] = "yinka"
	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // slices
	// var scores = []int{100,90,80}
	// fmt.Println(scores)
	// scores[2] = 34
	// scores = append(scores, 67)
	// fmt.Println(scores)

	// fmt.Print("Enter your number")
	// var input float64
	// fmt.Scanf("%f", &input)

	// output := input * 2
	// fmt.Println(output) 

	fmt.Print("Enter Fahrenheit value")
	var temp float64
	fmt.Scanln(&temp)

	ans := (temp * 9/5) + 32
	fmt.Println(ans) 

	fmt.Print("Enter feet")
	var feet float64
	fmt.Scanln(&feet)

	meters := feet * 0.3048
	fmt.Println(meters)
}