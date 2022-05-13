package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // "-" means this field are not reflected who are using this api
	Tags     []string `json:"tags,omitempty"` // "omitempty" says that if the valueof field is null or nil then dont throw that field at all
}

func main() {

	fmt.Println("This is code for create json data")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {

	ankCourses := []course{

		{"ReactJs", 299, "ankushweb", "ankush", []string{"Web-dev", "js"}},
		{"MongoDB", 199, "ankushweb", "ankush", nil},
		{"NodeJs", 188, "ankushweb", "ankush", []string{"Web-dev", "js"}},
	}

	// package this data as json data

	//finalJson, err := json.Marshal(ankCourses)
	finalJson, err := json.MarshalIndent(ankCourses, "", "\t")

	if err != nil {

		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	
		{
			"coursename": "NodeJs",
			"Price": 188,
			"website": "ankushweb",
			"tags": ["Web-dev","js"]
		}
	`)

	var ankCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {

		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &ankCourses)
		fmt.Printf("%#v\n", ankCourses)
	} else {

		fmt.Println("Json was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {

		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
