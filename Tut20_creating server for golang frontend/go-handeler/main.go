package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("This is handeler")
	//PerformGetRequest()
	PerformPostJsonRequest()
	//PerformPostFormRequest()

}

func PerformGetRequest() {

	const myurl = "http://localhost:4000"

	response, err := http.Get(myurl)

	if err != nil {

		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(responseString.String())

	fmt.Println(content)
	fmt.Println(string(content))

}

func PerformPostJsonRequest() {

	const myurl = "http://localhost:4000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"Go-lang",
			"price":0
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {

		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	// if err != nil {

	// 	panic(err)
	// }

	fmt.Println(string(content))

}

func PerformPostFormRequest() {

	const myurl = "http://localhost:4000/postform"

	data := url.Values{}
	data.Add("firstnname", "Ankush")
	data.Add("lastname", "Mukherjee")

	response, err := http.PostForm(myurl, data)

	if err != nil {

		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	response.Body.Close()

}
