package main

import "fmt"

// it is available globally
const loginToken string = "ankushislearninggolang" // Here capital L signifies the public variable

func main() {

	var username string = "Ankush"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n ", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n ", isLoggedIn)

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("variable is of type: %T \n ", smallvalue)

	// 	uint8       the set of all unsigned  8-bit integers (0 to 255) 2^8-1 including 0
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127) 2^7
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	var smallfloat float32 = 255.455
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n ", smallfloat)

	// float32     the set of all IEEE-754 32-bit floating-point numbers

	var smallfloat64 float64 = 255.455
	fmt.Println(smallfloat64)
	fmt.Printf("variable is of type: %T \n ", smallfloat64)

	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	// ######Default values or sum aliases

	// var anothervariable int
	// fmt.Println(anothervariable)
	// fmt.Printf("variable is of type: %T \n ", anothervariable)

	// impliclit type

	var website = "ankush.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n ", website)

	// No var style

	numberOfUser := 30000.0 // := it is not allowed in outside of main func or globally
	fmt.Println(numberOfUser)
	fmt.Printf("variable is of type: %T \n ", numberOfUser)
	numberOfUser = 64

	fmt.Println(loginToken)
	fmt.Printf("variable is of type: %T \n ", loginToken)

	// var name string
	// fmt.Println("enter your name ")
	// fmt.Scanf("%v", &name)
	// fmt.Println(name)

}
