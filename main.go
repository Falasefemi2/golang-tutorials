package main

import "fmt"

func highestScorer(scores map[string]int) string {
	// Your code here
	var highestScore int
	var highestScorer string

	for name, score := range scores {
		if score > highestScore {
			highestScore = score
			highestScorer = name
		}
	}

	return highestScorer


}

func main() {
	scores := map[string]int{
		"Alice": 85,
		"Bob": 92,
		"Charlie": 78,
		"David": 95,
	}
	fmt.Println(highestScorer(scores)) // Should print: "David"

	
}