package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	name := "isaac"
	other := "brian"

	fmt.Printf("Hello, %s and %s", name, other)

	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}
}
