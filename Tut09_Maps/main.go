package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps is golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["Rb"] = "Ruby"
	languages["py"] = "python"

	fmt.Println("list of all languages: ", languages)
	fmt.Println("js stands for: ", languages["js"])

	delete(languages, "Rb")
	fmt.Println("list of all languages: ", languages)

	// loops in golang

	for key, value := range languages {

		fmt.Printf("For key %v, value is %v \n", key, value)
	}

	for _, value := range languages {

		fmt.Printf("For key v, value is %v \n", value)
	}

}
