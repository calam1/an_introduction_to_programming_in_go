package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println(x) //x is still 5 like in java

	y := 5
	zeroPointer(&y)
	fmt.Println(y)

	z := 1.5
	square(&z)
	fmt.Println(z)

	a := 1
	b := 2

	swap(&a, &b)
	fmt.Println("swap a was 1 should be 2 and the swap result is", a)
	fmt.Println("swap b was 2 should be 1 and the swap result is", b)
}

func zero(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
