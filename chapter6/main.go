package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	averages()
	slices()
	slicesCopy()
	maps()
	mapOfMaps()
	smallestNumber()
}

func averages() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64
	//	for i := 0; i < len(x); i++ {
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
	fmt.Println("the 3rd element is: ", x[2])
	fmt.Println("values of 2:5 of x:", x[2:5])
}

func slices() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slices()", slice1, slice2)
}

func slicesCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("slicesCopy()", slice1, slice2)
}

func maps() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)
}

func mapOfMaps() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
	}

	if elm, ok := elements["Li"]; ok {
		fmt.Println(elm["name"], elm["state"])
	}
}

func smallestNumber() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	fmt.Println("max value of integer is:", MaxInt)

	min := MaxInt
	for _, values := range x {
		if values < min {
			min = values
		}
	}

	fmt.Println("the smallest number in the array is:", min)
}
