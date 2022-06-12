package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses-file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {

	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}
func main() {

	fmt.Println("server is running")

	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "Node js", CoursePrice: 299, Author: &Author{Fullname: "Ankush Mukherjee", Website: "Ankush web"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "React Js", CoursePrice: 175, Author: &Author{Fullname: "Sourav Nag", Website: "sourav web"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Javascript", CoursePrice: 160, Author: &Author{Fullname: "shyam", Website: "shyam web"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/getallcourses", getAllCourses).Methods("GET")
	r.HandleFunc("/getonecourse/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/createonecourse", createOneCourse).Methods("POST")
	r.HandleFunc("/updateonecourse/{id}", updateOneCourses).Methods("PUT")
	r.HandleFunc("/deleteonecourse/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to API by Ankush</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through courses, find matching id and return the response

	for _, course := range courses {

		if course.CourseId == params["id"] {

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	//return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")

	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {

		json.NewEncoder(w).Encode("please send some data")
	}

	// what about - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {

		json.NewEncoder(w).Encode("no data inside JSON")
	}

	// generate using id
	//	append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	//return

}

func updateOneCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one courses")
	w.Header().Set("Content-Type", "application/json")

	// grab id
	params := mux.Vars(r)

	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete One course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			break
		}

	}
}
