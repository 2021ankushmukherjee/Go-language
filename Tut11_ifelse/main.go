package main

import "fmt"

func main() {

	fmt.Println("if else in golang")

	loginCount := 5

	var result string

	if loginCount > 10 {

		result = "Regular User"
	} else if loginCount == 10 {

		result = "You are average user"
	} else {

		result = "You are not a user"
	}

	fmt.Println(result)

	if 9%2 == 0 {

		fmt.Println("Number is odd")
	} else {

		fmt.Println("Number is even")
	}

	if num := 3; num < 10 {

		fmt.Println("Num is less than 10")
	} else {

		fmt.Println(("Num is not less than 10"))
	}

}
