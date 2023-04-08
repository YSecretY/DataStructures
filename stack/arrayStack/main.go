package main

import "fmt"

const (
	maxSize = 10
)

type stack struct {
	top      int
	capacity int
	array    []int
}

func (s *stack) Create() *stack {
	s.capacity = maxSize
	s.top = -1
	s.array = make([]int, s.capacity)
	if s.array == nil {
		return nil
	}
	return s
}

func (s *stack) Empty() bool {
	return s.top == -1
}

func (s *stack) Full() bool {
	return s.top == s.capacity-2
}

func (s *stack) Push(data int) {
	if s.Full() {
		fmt.Println("stack overflow")
		return
	}
	s.top++
	s.array[s.top] = data
}

func (s *stack) Pop() {
	if s.Empty() {
		fmt.Println("stack is already empty")
		return
	}
	s.top--
}

func main() {
	s := stack{}
	s.Create()
	s.Push(5)
	s.Push(3)
	s.Push(8)
	s.Push(6)
	s.Push(4)
	s.Push(7)

	fmt.Printf("top of the stack is: %d\n", s.array[s.top])
	s.Pop()
	fmt.Printf("top of the stack after pop: %d\n", s.array[s.top])
	fmt.Printf("stack is empty: %t\n", s.Empty())
	fmt.Printf("stack is full: %t\n", s.Full())

	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(9)

	fmt.Printf("stack is full after adding elements: %t\n", s.Full())
}
