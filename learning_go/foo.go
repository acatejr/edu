package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "A text string"
		fmt.Println(y)
	}
}
