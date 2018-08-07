package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type DumbLinkedList struct {
  length int
  Head *Node
}

func (l *DumbLinkedList) add_front(n *Node){
  n.Next = l.Head
  l.Head = n
  l.length++
}

func (l *DumbLinkedList) remove_front() *Node{
  tmp := l.Head
  l.Head = l.Head.Next
  l.length--
  return tmp
}

func (l *DumbLinkedList) add_back(n *Node){
  current_pointer := l.Head
  for (current_pointer!=nil){
    if current_pointer.Next!=nil{
      current_pointer=current_pointer.Next
    } else {break}
  }
  current_pointer.Next = n
}

func (l *DumbLinkedList) remove_back() *Node{
  current_pointer := l.Head
  previous_pointer := current_pointer
  for (current_pointer!=nil){
    if current_pointer.Next!=nil{
      previous_pointer=current_pointer
      current_pointer=current_pointer.Next
    } else {break}
  }
  previous_pointer.Next=nil
  return current_pointer
}

func (l DumbLinkedList) print_list(){
  current_pointer := l.Head
  for (current_pointer!=nil){
    fmt.Println(*current_pointer)
    current_pointer=current_pointer.Next
  }
  fmt.Println("//-----//")
}

func main() {
  n1 := Node{1, nil}
  n2 := Node{2, nil}
  n3 := Node{3, nil}
  dumb_list := DumbLinkedList{0, nil}
  dumb_list.add_front(&n1)
  dumb_list.add_front(&n2)
  dumb_list.add_back(&n3)
  dumb_list.print_list()
  dumb_list.remove_front()
  dumb_list.print_list()
  dumb_list.remove_back()
  dumb_list.print_list()
}
