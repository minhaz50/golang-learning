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

/*
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

*/
	func getWelcome() {
		fmt.Println("Welcome to the application.")
	}

	func getNameInput() string {
		var name string
		fmt.Println("Enter Your name: ")
		fmt.Scanln(&name)
		return name
	}

	func getNumberInput() (int, int){
		var num1 int
		var num2 int

		fmt.Println("Enter Your first number: ")
		fmt.Scanln(&num1)

		fmt.Println("Enter Your Second number: ")
		fmt.Scanln(&num2)

		return num1 , num2
	}

	func addNumber(num1 int  , num2 int) int {
		sum := num1 + num2
		return  sum
	}

	func displayMessage(name string, sum int){
		fmt.Println("Hello, ", name)
		fmt.Println("Your Summation of the two numbers are, ", sum)
	}

	func goodByeMessage() {
		fmt.Println("Thank you for using this application.")
		fmt.Println("Goodbye")
	}

	func main() {		
		getWelcome() // welcome message
		name := getNameInput() // Name Input
		num1, num2 := getNumberInput() // Number Input
		sum := addNumber(num1, num2)  // Add Number 
		displayMessage(name, sum) // display result 
		goodByeMessage() // Goodbye Message
		

	}