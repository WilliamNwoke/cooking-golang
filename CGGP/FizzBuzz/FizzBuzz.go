package main

import (
	"fmt"
)

func main() {
	//  scanner object to read from Standard input

	for num := 1; num <= 20; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if num%3 == 0 {
				fmt.Println("Fizz")
			} else if num%5 == 0 {
				fmt.Println("Buzz")
			}
		}
	}

}
