package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func main() {
	v := Node{1, nil}
	p := &v
	p.Value = 1e9
	fmt.Println(v)
	head := Node{0, p}
	fmt.Println(*head.Next)
}
