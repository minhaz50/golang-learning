package main

import "fmt"

/*
	Array Declaration
	var arrayName [size]Type

	exm:
		var marks [3]int
		marks[0] = 90
		marks[1] = 80
		marks[2] = 70

	There are different way to declare array

	initialize together
		numbers := [4]int{1,2,3,4}

		nums := [5]int{1,2}
		if we print this output will be [1,2,0,0,0]

*/


func main () {
	/*
	numbers := [3]int{10,20,30}

	numbers[2] = 40 // it's change the value


	for i :=0; i <len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println(len(numbers)) // len() give the length of the array 
	fmt.Println(numbers[0])  // access element
	fmt.Println(numbers[2]) 
	*/

	nums := [5] int {10, 20, 30, 40, 50}

	fmt.Println("Print All numbers:", nums)
	fmt.Println("Length", len(nums))

	fmt.Println("Accessing...")

	for i,v := range nums {
		fmt.Println(i,v)
	} 
}