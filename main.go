package main

import (
	"fmt"
)  

func main() {
	//create a variable with type int(integer)
	var n int

	fmt.Println("WELCOME TO DECTOBIN")
	fmt.Println("Pick a number to convert:")

	fmt.Scan(&n) //This code allows you to input a number 
	
	fmt.Printf("%b", n)  // This code converts the integer to binary
}
