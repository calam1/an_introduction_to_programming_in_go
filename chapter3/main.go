package main

import "fmt"

func main() {
	fmt.Println("0 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x string
	x = "Hello world"
	fmt.Println(x)

	x = "Goodbye world"
	fmt.Println(x)

	x += " some other saying"
	fmt.Println(x)

	var y string = "hello"
	var z string = "hello"
	//this is true
	fmt.Println(y == z)

	z = "there"
	//this is false
	fmt.Println(y == z)

	a := "chris"
	fmt.Println(a)

	b := 5
	fmt.Println(b)

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

}
