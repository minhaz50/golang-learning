package main

import "fmt"

func main() {
	var name string = "Minhaz"
	var isGoFun bool = true
	
	fmt.Println("My name is: ",name,".","Yes Go is Fun:",isGoFun)

	// Print
	var x , y string = "Hello", "World"
	fmt.Print(x,y,"\n")
	
	//  Println,
	var a, b string = "Hello", "World"
	println(a,b)


	// Printf 
	var i, j string  = "Hello", "world"
	fmt.Printf("i has value %v and type %T\n", i, i)
	fmt.Printf("j has value %v and type %T", j, j)


	fmt.Println("\n-------------------------------------------")
	var s = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", s)
  fmt.Printf("%#v\n", s)
  fmt.Printf("%v%%\n", s)
  fmt.Printf("%T\n", s)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)

}