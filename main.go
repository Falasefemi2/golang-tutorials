package main

import "fmt"

func main()  {

	age := 24
	name := "David"
	// Print
	fmt.Print("Femi, ")
	fmt.Print("Falase \n")
	fmt.Print("Samuel \n")

	fmt.Println("Hello World!")
	fmt.Println("Goodbye World!")

	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted string) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v", age, name)

}