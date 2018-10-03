package main

import (
	"fmt"
)

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world!")
}

func main() {

	student := []string{}
	students := make([]string, 35, 50)
	moreStudents := []string{"Mike", "Bill"}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(moreStudents)

	for i := 1; i < 1000; i++ {
		moreStudents = append(moreStudents, "Steve")
	}

	fmt.Println(moreStudents)
	// defer world()
	// hello()
	// x := 42
	// fmt.Println(x)
	// {
	// 	fmt.Println(x)
	// 	y := "A text string"
	// 	fmt.Println(y)
	// }
}
