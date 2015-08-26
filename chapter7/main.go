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

	factorialSum := factorial(4)
	fmt.Println("factorial sum", factorialSum)

	defer second()
	first()

	//	defer func() {
	//		str := recover()
	//		fmt.Println(str)
	//	}()
	//	panic("PANIC")

	sumArray := []int{1, 2, 3, 4, 5}
	sumVal := sum(sumArray)
	fmt.Println("sum", sumVal)

	fmt.Println(oddEvenHalf(2))
	fmt.Println(oddEvenHalf(5))

	fmt.Println(maxValueInVariadic(1, 6, 2, 4, 9, 0))

	fmt.Println("fib 0", fib(0))
	fmt.Println("fib 1", fib(1))
	fmt.Println("fib 2", fib(2))
	fmt.Println("fib 3", fib(3))
	fmt.Println("fib 4", fib(4))
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

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func sum(x []int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func oddEvenHalf(x int) (int, bool) {
	var evenOdd bool

	if x%2 == 0 {
		evenOdd = true
	}

	return x / 2, evenOdd
}

func maxValueInVariadic(x ...int) int {
	max := 0

	for _, val := range x {
		if val > max {
			max = val
		}
	}

	return max
}

func fib(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	return fib(x-2) + fib(x-1)
}
