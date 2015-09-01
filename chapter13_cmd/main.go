package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//define flags
	maxp := flag.Int("max", 6, "the max value")
	//parse
	flag.Parse()
	//seed to not return the same number
	rand.Seed(time.Now().Unix())
	//generate a number between 0 and maxp
	fmt.Println(rand.Intn(*maxp))
}
