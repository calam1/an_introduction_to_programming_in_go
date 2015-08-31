package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	//createFile("this is a golang test")
	//readFile("testclam.txt")
	//readDirectory()
	readDirectoryRecursively()

	//container/list example
	var llist list.List
	llist.PushBack(1)
	llist.PushBack(2)
	llist.PushBack(3)

	for e := llist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []Person{
		{"Chris", 9},
		{"Sue", 9},
		{"Brad", 9},
		{"Jill", 9},
		{"Jack", 10},
	}
	fmt.Println(kids)
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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
