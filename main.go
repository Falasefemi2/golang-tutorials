package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	name := [4]string{"femi","sameul","ayo","dare"}
	name[1] = "yinka"

	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))

	// slices (use arrays under the hood)
	var scores = []int{100,90,30}
	fmt.Println(scores)
	scores[2] = 60
	fmt.Println(scores)

	scores = append(scores, 87)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := name[1:4]
	rangeTwo := name[2:]
	rangwThree := name[:3]
	fmt.Println(rangeOne, rangeTwo,rangwThree)
}