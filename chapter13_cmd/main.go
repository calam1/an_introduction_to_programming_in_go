package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	//define flags
	maxp := flag.Int("max", 6, "the max value")
	//parse
	flag.Parse()
	//generate a number between 0 and maxp
	fmt.Println(rand.Intn(*maxp))
}
