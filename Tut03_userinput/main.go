package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give the rating of our pizza: ")

	// comma ok // error ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of this rating is: %T ", input)
	fmt.Println("")

	// input, error := reader.ReadString('\n')
	// fmt.Println("Thanks for rating: ", input)
	// fmt.Printf("Type of this rating is: %T ", input)
	// fmt.Println("")
	// fmt.Println(error)

}
