package main

import (
	"container/list"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
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

	//sort example
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

	//hash example
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	//hash file compare example
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}

	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)

	//crypto hash
	cryptosha := sha1.New()
	cryptosha.Write([]byte("test"))
	bs := cryptosha.Sum([]byte{})
	fmt.Println(bs)

	//client server example
	//go server()
	//go client()

	//var input string
	//fmt.Scanln(&input)

	//http server
	//http.HandleFunc("/hello", hello)
	//http.ListenAndServe(":9000", nil)

	//rpc server example
	go rpcServer()
	go rpcClient()
	var input string
	fmt.Scanln(&input)
}

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func rpcServer() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func rpcClient() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				Hello World!
			</body>
		</html>`,
	)
}

func server() {
	//listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	//connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
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
