package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}

		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
		default:
			fmt.Println("unknown number")
		}
		i += 1
	}

	x := 10
	if x > 10 {
		fmt.Println("big")
	} else {
		fmt.Println("small")
	}

	fizzfuzz()
}

func fizzfuzz() {
	i := 1

	for i <= 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}

		i += 1
	}
}
