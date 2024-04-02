package main

import "fmt"



func main() {
xs := []float64{1,2,3,4,5,6,7,8,9,10}

total := 0.0
for _, v := range xs {
	total += v
}
fmt.Println(total / float64(len(xs)))
}
