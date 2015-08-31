package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//createFile("this is a golang test")
	readFile("testclam.txt")
}

func createFile(txt string) {
	file, err := os.Create("testclam.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(txt)
}

func readFile(file string) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
