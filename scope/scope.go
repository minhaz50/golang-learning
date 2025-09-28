package main

import "fmt"

var a int = 20
var b int = 30

func add(x int, y int)  {
	z := x + y
	fmt.Println(z)
}


func main(){
	fmt.Println("Adding Number")
	p := 50
	q := 60

	add(p,q)

	add(a,b)

	add(p, a)

	add (p, b)
day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default: 
	fmt.Println("Another day")
	}
 

}