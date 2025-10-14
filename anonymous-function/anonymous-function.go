package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}

	result := add(5, 5)
	fmt.Println(result)
	
	counter := 0

	increment := func () int {
		counter++
		return counter
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}