package main

import "fmt"

func main() {

	defer fmt.Println("Hello world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("This is Defer")
	myDefer()

}

func myDefer() {

	for i := 0; i < 5; i++ {

		defer fmt.Print(i, "\n")

	}
}
