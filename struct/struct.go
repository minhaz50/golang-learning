package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{
		Name: "Minhaz",
		Age:  28,
	}

	fmt.Println("Name:" , user1.Name)
	fmt.Println("Age: ", user1.Age)

	user2 := User{
		Name: "Israt",
		Age: 21,
	}

	fmt.Println("Name:" , user2.Name)
	fmt.Println("Age: ", user2.Age)
}