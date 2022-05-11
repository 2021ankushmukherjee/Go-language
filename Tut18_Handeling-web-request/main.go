package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const url = "https://lco.dev"

const url = "https://www.google.com/"

func main() {

	fmt.Println("This is web request handeling")

	response, err := http.Get(url)

	if err != nil {

		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {

		panic(err)
	}

	//fmt.Println(databytes)
	content := string(databytes)
	fmt.Println(content)
	fmt.Println("Status code is", +response.StatusCode)

}
