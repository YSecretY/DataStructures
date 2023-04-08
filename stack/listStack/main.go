package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		fmt.Println("stack is empty")
		return 0
	}
	return s.head.data
}

func (s *Stack) Push(data ...interface{}) {
	for _, val := range data {
		s.head = &Node{val, s.head}
		s.size++
	}
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.Empty() {
		fmt.Println("stack is already empty")
		return 0, false
	}
	data := s.head.data
	s.head = s.head.next
	s.size--
	return data, true
}

func (s *Stack) Print() {
	for cur := s.head; cur != nil; cur = cur.next {
		fmt.Print(cur.data, " ")
	}
	fmt.Printf("\n")
}

func main() {
	s := Stack{}

	s.Push(6, 4, 7, 2, 4, 8, 9, 15)
	fmt.Print("stack after push: ")
	s.Print()

	s.Pop()
	fmt.Print("stack after pop: ")
	s.Print()

	fmt.Printf("stack top: %v\n", s.Top())
	fmt.Println("stack is empty:", s.Empty())
	fmt.Println("stack size is:", s.Size())
}
