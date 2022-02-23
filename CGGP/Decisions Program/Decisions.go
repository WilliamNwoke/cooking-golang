//
// Decisions program from the bootcamp 1 in 2022, as taught by professor Mike Mckee.
// This is a golang version of the assignments completed in java
//
// @author Uchenna Nwoke
// @version 1.0
//
package main

import (
	"fmt"
)

// The decision program uses an if statement to compare two integers, and checks
// if the first integer number is a multiple of the second, with its factor or
// not.
func main() {
	// Welcome message to the program
	fmt.Println("Welcome to the decisions program!")

	// Prompts the user to enter two values seperated by a space
	fmt.Println("\nEnter two integer numbers, seperated by a space: ")

	// integer variable declarations for the for the first and second number
	var firstNum, secondNum int

	// fmt.scanf scans for the first and second number passed in standard in
	// numbers must be seperated by a space for this program to work
	fmt.Scanf("%d %d", &firstNum, &secondNum)

	// Decison Statement to check if the first number is a multiple of the second number
	// and prints out the corresponding message
	if (firstNum % secondNum) == 0 {
		factor := firstNum / secondNum
		fmt.Printf("%d is a multiple of %d with a factor of %d\n",
			firstNum, secondNum, factor)
	} else {
		fmt.Printf("%d is not a multiple of %d\n", firstNum, secondNum)
	}

	// Goodbye message
	fmt.Println("Thanks for using the decision program")

}
