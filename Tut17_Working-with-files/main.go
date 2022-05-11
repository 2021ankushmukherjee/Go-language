package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("This is files code")
	content := "This needs to go in a file - Ankush Wants to be a good developer"

	file, err := os.Create("./mylocalfile.txt")

	if err != nil {

		panic(err)

	}

	length, err := io.WriteString(file, content)

	if err != nil {

		panic(err)

	}

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./mylocalfile.txt")
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)

	// if err != nil {

	// 	panic(err)
	// }

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {

	if err != nil {

		panic(err)
	}
}
