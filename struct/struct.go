package main

import "fmt"

type User struct {
	Name string  // this is called member variable or property
	Age  int
}

func main() {
	user1 := User{
		Name: "Minhaz",
		Age:  28,
		
	}
	user2 := User{
		Name: "Israt",
		Age: 21,
	}

	user1.Age = 30

	fmt.Println("Name:" , user1.Name)
	fmt.Println("Age: ", user1.Age)
	fmt.Println("Name:" , user2.Name)
	fmt.Println("Age: ", user2.Age)
}