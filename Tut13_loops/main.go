package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {

	// 	fmt.Println(days[d])
	// }

	// for i := range days {

	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {

	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {

			goto ankush
		}

		// if rougeValue == 5 {

		// 	break
		// }

		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

	// goto

ankush: // this is label
	fmt.Println("This is goto")
}
