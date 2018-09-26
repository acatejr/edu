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

	defer world()
	hello()
	// x := 42
	// fmt.Println(x)
	// {
	// 	fmt.Println(x)
	// 	y := "A text string"
	// 	fmt.Println(y)
	// }
}
