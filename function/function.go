package main

import "fmt"

/*
func add (num1 int, num2 int){
	sum := num1 + num2

	fmt.Println(sum)
}


func main() {
	a := 10
	b := 10

	add(a, b)
	add(3,7)
}
*/

	func add (num1 int, num2 int)( int, int){
		sum := num1 + num2 
		multiplication := num1 * num2

		return sum, multiplication
	}

	func main( ){
		a := 10
		b := 20

		p, q:= add(a,b)
		fmt.Println(p, q)
	}