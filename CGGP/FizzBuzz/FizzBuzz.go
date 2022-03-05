package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//  scanner object to read from Standard input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nEnter an integer number: ")
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 64)

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
