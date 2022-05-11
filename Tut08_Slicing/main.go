package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slicing")

	var fruitList = []string{"Apple", "Mango", "Grapes"}
	fmt.Printf("The type of fruitList %T\n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Banana", "Kiwi")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // start from element 1
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // start from 1 and go upto 3
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3]) // start from 1 and go upto 3
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	
	highScores[0] = 234
	highScores[1] = 264
	highScores[2] = 365
	highScores[3] = 655

	highScores = append(highScores, 555, 666, 777) // it reallocates the memory and all of the values are acomodated.

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) // gives a boolean value
	sort.Ints(highScores)
	fmt.Println(highScores)

	// remove a value from slices based on index

	var courses = []string{"reactjs", "node js", "javascript", "swift", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //ignore element
	fmt.Println(courses)

}
