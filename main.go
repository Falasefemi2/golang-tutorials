package main

import "fmt"

func half(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true // return half of n and true if n is even
	} else {
		return 0, false // return 0 and false if n is odd
	}
}

func main() {
	result, isEven := half(6)
	fmt.Println(result, isEven) // Output: 0 false
}
