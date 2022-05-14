package main

import "fmt"

func main() {

	fmt.Println("This is a pointer")
	var ptr *int // create a pointer

	fmt.Println("value of pointer is", ptr)

	mynumber := 24

	mynumber = mynumber + 2
	fmt.Println(mynumber)

	var ptr1 = &mynumber // reference to memory location
	fmt.Println("Value of memory location is ", ptr1)
	fmt.Println("Value of actual pointer is ", *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println("The value is: ", mynumber)

}
