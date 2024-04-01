package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n )
} 

func sayBye(n string) {
	fmt.Printf("GoodBye %v \n", n )
} 

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circlArea (r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Femi")
	// sayGreeting("Samuel")
	// sayBye("Dare")
	cycleNames([]string{"femi","sam","ayo"}, sayGreeting)
	cycleNames([]string{"femi","sam","ayo"}, sayBye)
	
	a1 := circlArea(10.5)
	a2 := circlArea(15.5)
	fmt.Println(a1,a2)
	fmt.Printf("circle 1 is %0.3f  and circle 2 is %0.3f",a1, a2)
}


