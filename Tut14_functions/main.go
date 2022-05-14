package main

import "fmt"

func main() {

	fmt.Println("This is functions")
	greeter()
	result := adder(3, 5)
	fmt.Println(result)

	prores := proAdder(2, 4, 6, 8)
	fmt.Println(prores)
}

func greeter() {

	fmt.Println("Good morning")
}

func adder(x int, y int) int {

	return x + y
}

func proAdder(values ...int) int {

	total := 0

	for _, val := range values {

		total += val
	}

	return total
}
