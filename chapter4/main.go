package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	//	output := input * 2
	output := (input - 32) * 5 / 9

	fmt.Println(output)
}
