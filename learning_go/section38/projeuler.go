package main

import "fmt"

func main() {
	below := 5
	mults := []int{2, 3}
	fmt.Printf("The sum of 1, 2, 3, 4, 5 using multiples of 2 and 3 is: %d\n", sumOfMultiplesBelow(below, mults))

	below = 10
	mults = []int{3, 5}
	fmt.Printf("The sum of 1, 2, 3, 4, 5 using multiples of 2 and 3 is: %d\n", sumOfMultiplesBelow(below, mults))

	below = 1000
	mults = []int{3, 5}
	// 233168
	fmt.Printf("The sum of 1, 2, 3, 4, 5 using multiples of 2 and 3 is: %d\n", sumOfMultiplesBelow(below, mults))
}

func sumOfMultiplesBelow(below int, multipes []int) int {
	sum := 0
	for v := below - 1; v > 0; v-- {
		if v%multipes[0] == 0 || v%multipes[1] == 0 {
			sum += v
		}
	}
	return sum
}
