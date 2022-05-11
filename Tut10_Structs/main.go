package main

import "fmt"

func main() {

	fmt.Println("this is structs code") // structs stands for structure

	// no inheritance, super, parents in go lang

	ankush := User{"Ankush Mukherjee", "ankush@gmail.com", true, 22}

	fmt.Println(ankush)
	fmt.Printf("ankush detials are: %+v\n", ankush)
	fmt.Printf("Your name is %v and email is %v\n", ankush.Name, ankush.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

