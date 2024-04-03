package main

import "fmt"

func main() {
	mybill := newBill("femi falase's bill")

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
