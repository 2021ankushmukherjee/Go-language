package main

import "fmt"

func main() {

	fmt.Println("Welcome to arrays")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"
	fruitList[3] = "Grapes"
	fruitList[4] = "Orange"

	fmt.Println(fruitList)
	fmt.Println("Length of fruit list: ", len(fruitList))

	var veglist = [4]string{"Potato", "beans", "Brinjal", "Tomato"}
	fmt.Println(veglist)
	fmt.Printf("%T\n", veglist)
}
