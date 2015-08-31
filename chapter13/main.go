package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//createFile("this is a golang test")
	//readFile("testclam.txt")
	//readDirectory()
	readDirectoryRecursively()
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

func readDirectory() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func readDirectoryRecursively() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
