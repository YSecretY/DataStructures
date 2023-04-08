package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) First() interface{} {
	return q.front.data
}

func (q *Queue) Last() interface{} {
	return q.rear.data
}

func (q *Queue) Empty() bool {
	return q.front == nil
}

func (q *Queue) Append(data ...interface{}) {
	for _, val := range data {
		cur := &Node{val, nil}
		if q.front == nil {
			q.front = cur
			q.rear = cur
		} else {
			q.rear.next = cur
			q.rear = cur
		}
	}
}

func (q *Queue) Dequeue() {
	if q.Empty() {
		fmt.Println("queue is already empty")
		return
	}
	q.front = q.front.next
}

func main() {
	q := Queue{}
	q.Append(10, 20, 30, 40, 50)

	fmt.Printf("first element to out is: %v\n", q.First())
	fmt.Printf("last element to out is: %v\n", q.Last())
	fmt.Println("queue is empty:", q.Empty())

}
