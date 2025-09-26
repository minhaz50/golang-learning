package main

import "fmt"

func main () {

	// if-else 
	age := 17.9;
	if age > 18 {
		fmt.Println("You are eligible for Voter")
	}else if age < 18{
		fmt.Println("You are not eligible for Voter")
	} else if age == 18{
		fmt.Println("You are just a teenager, not eligible to be a voter.")
	}


	// switch  

	day := 0

	switch day{
	case 1: 
		fmt.Println("Sunday")
	
	case 2: 
	fmt.Println("Monday")
		
	case 3:
		fmt.Println("Tuesday")
	
	default: 
		fmt.Println("Another day")

	}
}