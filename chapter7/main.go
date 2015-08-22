package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	x, y := manyValues()
	fmt.Println(x, y)

	fmt.Println(add(1, 2, 3, 4))

	add2 := func(x, y int) int {
		return x + y
	}

	fmt.Println(add2(2, 4))

	z := 0
	increment := func() int {
		z++
		return z
	}

	fmt.Println(increment())
	fmt.Println(increment())

	factorial(4)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, x := range xs {
		total += x
	}

	avg := total / float64(len(xs))

	return avg
}

func manyValues() (int, int) {
	return 6, 7
}

func add(nums ...int) int {
	total := 0
	for _, vals := range nums {
		total += vals
	}

	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
