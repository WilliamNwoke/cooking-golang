package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//  scanner object
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an adjective: ")
	scanner.Scan()
	adj := scanner.Text()

	fmt.Print("Enter a verb: ")
	scanner.Scan()
	verb1 := scanner.Text()

	fmt.Print("Enter another verb: ")
	scanner.Scan()
	verb2 := scanner.Text()

	fmt.Print("Enter a famous person: ")
	scanner.Scan()
	famousPerson := scanner.Text()

	madlibs := fmt.Sprintf("Computer pramming is so %s! It makes me so excited all the time because \n"+
		" I love to %s. Stay hydrated and %s like you are %s", adj, verb1, verb2, famousPerson)

	fmt.Printf(madlibs)
}
