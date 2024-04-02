package main

import "fmt"

func updateName(x string)  {
	x = "wedge"
}


func main() {
	name := "tife"

	updateName(name)

	fmt.Println("memory address of name is : ", &name)

	m := &name
	fmt.Println(m)
	fmt.Println("value at memory addresss: ", *m)
	fmt.Println(name)


}
