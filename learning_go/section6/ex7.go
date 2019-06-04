package main

import "fmt"

// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of
// all the multiples of 3 or 5 below 1000
func main() {
	sum := 0
	fmt.Println("Running exercise 6, section 6.")
	for i := 999; i >= 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("The sum of all natural numbers that are muliples of 3 or 5 less than 1000 is: %d\n", sum)
	fmt.Println("Done!")
}
