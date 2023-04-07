package main

import (
	"fmt"
)

type el struct {
	data interface{}
	prev *el
	next *el
}

type list struct {
	pbeg *el
	pend *el
	len  int
}

func (l *list) begin() *el {
	return l.pbeg
}

func (l *list) end() *el {
	return l.pend.next
}

func (l *list) length() int {
	return l.len
}

func (l *list) append(data ...interface{}) {
	for _, val := range data {
		l.len++
		cur := &el{val, nil, nil}
		if l.pbeg == nil {
			l.pbeg = cur
		} else {
			l.pend.next = cur
			cur.prev = l.pend
		}
		l.pend = cur
	}
}

func (l *list) prepend(data ...interface{}) {
	for _, val := range data {
		l.len++
		cur := &el{val, nil, nil}
		if l.pbeg == nil {
			l.pbeg = cur
			l.pend = cur
		} else {
			l.pbeg.prev = cur
			cur.next = l.pbeg
		}
		l.pbeg = cur
	}
}

func (l *list) erase(el *el) *el {
	l.len--
	if el == l.pbeg && el == l.pend {
		el, l.pbeg, l.pend = nil, nil, nil
		return nil
	}

	if el == l.pbeg {
		l.pbeg = el.next
		l.pbeg.prev = nil
		return l.pbeg
	}

	if el == l.pend {
		l.pend = l.pend.prev
		l.pend.next = nil
		return nil
	}

	el.prev.next = el.next
	el.next.prev = el.prev
	return el.next
}

func (l *list) insert(data interface{}, before *el) *el {
	if before == nil {
		l.append(data)
		return l.pend
	}
	l.len++

	cur := &el{data, nil, nil}
	if l.len == 0 {
		l.pbeg = cur
		l.pend = cur
		return cur
	}

	if l.len == 1 {
		l.pbeg.prev = cur
		cur.next = l.pbeg
		l.pend = l.pbeg
		l.pbeg = cur
		return cur
	}

	if before == l.pbeg {
		l.pbeg.prev = cur
		cur.next = l.pbeg
		l.pbeg = cur
		return cur
	}

	if before == l.pend {
		l.pend.prev.next = cur
		l.pend.prev = cur
		cur.next = l.pend
		return cur
	}

	before.prev.next = cur
	before.prev = cur
	cur.next = before

	return cur
}

func (l *list) clear() {
	for cur := l.begin().next; cur != l.end(); cur = cur.next {
		cur.prev.next = nil
		cur.prev = nil
	}
	l.pbeg = nil
	l.len = 0
}

func (l *list) print() {
	cur := l.begin()
	for cur != l.end() {
		fmt.Printf("%v ", cur.data)
		cur = cur.next
	}
	fmt.Printf("\n")
}

func main() {
	l := list{}
	l.append(18, 25, 132, 4, 2, 10, 295, 169)
	fmt.Print("list before deleting elements: ")
	fmt.Println("length =", l.length())
	l.print()

	for cur := l.begin(); cur != l.end(); cur = cur.next {
		if cur.data == 18 || cur.data == 25 || cur.data == 2 || cur.data == 169 {
			l.erase(cur)
		}
	}

	fmt.Print("list after deleting elements: ")
	fmt.Println("length =", l.length())
	l.print()

	fmt.Print("list after clearing elements: ")
	l.clear()
	fmt.Println("length =", l.length())
	l.print()

	fmt.Print("list after prepending elements: ")
	l.prepend(1, 2, 3, 4, 5)
	fmt.Println("length =", l.length())
	l.print()

	for cur := l.begin(); cur != l.end(); cur = cur.next {
		if cur.data == 5 {
			l.insert(6, cur)
		}
		if cur.data == 3 {
			l.insert(3, cur)
		}
		if cur.data == 1 {
			l.insert(12, cur)
		}
	}

	fmt.Print("list after inserting elements: ")
	fmt.Println("length =", l.length())
	l.print()
}
