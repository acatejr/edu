package main

import "fmt"

func main() {
	fmt.Println("Excercise 1.")
	fmt.Println(half(1))
	fmt.Println(half(2))

	// Excercise 2
	doHalf := func(n int) (int, bool) {
		isEven := (n%2 == 0)
		result := n / 2
		return result, isEven
	}
	fmt.Println("")
	fmt.Println(doHalf(1))
	fmt.Println(doHalf(2))
	fmt.Println(doHalf(5))

	// Excercise 3 - find the greatest integer in list
	fmt.Println("")
	fmt.Printf("The max is %d\n", findGreatest(1, 2, 3))
	fmt.Printf("The max is %d\n", findGreatest(3, 99, 3))
	fmt.Printf("The max is %d\n", findGreatest(13, 459, 7, 2, 4, 53, 73, 25, 3))

	// Excercise 4 what is the value of:
	// (true && false) || (false && true) || !(false && fale)
	fmt.Println("")
	fmt.Printf("%t", expressionTest())

	// Excercise 5 - function foo called in many ways
	fmt.Println("")
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

}

func foo(args ...int) {
	fmt.Println(args)
}

func expressionTest() bool {
	return (true && false) || (false && true) || !(false && false)
}

func findGreatest(numList ...int) int {
	var result int
	for i := range numList {
		if numList[i] > result {
			result = numList[i]
		}
	}
	return result
}

// Excercise 1
func half(n int) (int, bool) {
	isEven := (n%2 == 0)
	result := n / 2
	return result, isEven
}
