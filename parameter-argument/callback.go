package main

import "fmt"

func processData(callback func(string)) {
	data := "User data loaded"
	callback(data)
}

func main() {
	processData(func(result string) {
		fmt.Println("callback received", result)
	})
}