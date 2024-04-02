package main

import "fmt"

func mergeInventories(store1, store2 map[string]int) map[string]int {
	// Your code here
	merged := make(map[string]int)
	for item, quantity := range store1 {
		merged[item] += quantity
	}
	for item, quantity := range store2 {
		merged[item] += quantity
	}

	return merged

}

func main() {
	store1 := map[string]int{
		"apple":  10,
		"banana": 20,
	}

	store2 := map[string]int{
		"banana": 15,
		"orange": 25,
	}

	merged := mergeInventories(store1, store2)
	fmt.Println(merged) // Should print: map[apple:10 banana:35 orange:25]
}
