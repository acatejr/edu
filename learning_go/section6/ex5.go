package main

import "fmt"

// Print all of the even numbers between 0 and 100
func main() {
	fmt.Println("Running exercise 5, section 6.")
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Done!")
}
