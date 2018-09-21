package main

import "fmt"

// Create a program that prints to the terminal asking for
// a user to enter a small number and a larger number. Print
// the remainder of the larger number divided by the smaller number.
func main() {
	fmt.Println("Running exercise 4, section 6.")
	var smallNum int
	var largeNum int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&smallNum)
	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNum)
	fmt.Printf("The remainder is %d.\n", largeNum%smallNum)
	fmt.Println("Done!")
}
