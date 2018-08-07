package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type DumbSortedLinkedList struct { //Sorted from small to large
  length int
  Head *Node
}

func (l *DumbSortedLinkedList) add(n *Node){
	// fmt.Print("Adding ")
	// fmt.Println(*n)
	if l.Head == nil{
		l.Head = n
		l.length++
		return
	}

	if n.Value < l.Head.Value{
		n.Next = l.Head
	  l.Head = n
	  l.length++
		return
	} else {
		current_pointer := l.Head
		for (current_pointer.Next!=nil){
			if n.Value <= current_pointer.Next.Value{ //Add in middle
				tmp_pointer := current_pointer.Next
				current_pointer.Next=n
				current_pointer.Next.Next=tmp_pointer
				l.length++
				return
			}
			current_pointer=current_pointer.Next
	  }
		current_pointer.Next=n
		return
	}
}

func (l *DumbSortedLinkedList) remove_front() *Node{
  tmp := l.Head
  l.Head = l.Head.Next
  l.length--
  return tmp
}

func (l *DumbSortedLinkedList) remove_back() *Node{
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

func (l DumbSortedLinkedList) print_list(){
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
  dumb_list := DumbSortedLinkedList{0, nil}
  dumb_list.add(&n1)
  dumb_list.add(&n2)
  dumb_list.add(&n3)
  dumb_list.print_list()
	dumb_list.remove_front()
	dumb_list.print_list()
	dumb_list.remove_back()
	dumb_list.print_list()

}
