package main

import "fmt"

func totalInventoryValue(inventory map[string]int) int {
	// Your code here
	total := 0

	for _, v :=  range inventory {
		total += v
	}
	return total
}

func main() {
	inventory := map[string]int{
		"apple":  50,
		"banana": 30,
		"orange": 40,
	}
	
	fmt.Println(totalInventoryValue(inventory)) // Should print: 120

	
}