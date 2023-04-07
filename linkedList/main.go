package main

import "fmt"

type el struct {
	data interface{}
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
		cur := &el{val, nil}
		if l.pbeg == nil {
			l.pbeg = cur
			l.pend = cur
		}

		l.pend.next = cur
		l.pend = cur
	}
}

func (l *list) prepend(data ...interface{}) {
	for _, val := range data {
		l.len++
		cur := &el{val, l.pbeg}
		if l.pbeg == nil {
			l.pend = cur
		}
		l.pbeg = cur
	}
}

func (l *list) insert(data interface{}, after *el) *el {
	if after == nil {
		l.append(data)
		return l.pend
	}
	l.len++
	cur := &el{data, nil}

	if l.pend == after {
		l.pend.next = cur
		l.pend = cur
		return cur
	}

	if after == l.pbeg {
		cur.next = l.pbeg.next
		l.pbeg.next = cur
		return cur
	}

	cur.next = after.next
	after.next = cur

	return cur
}

func (l *list) erase(val *el) {
	if val == nil {
		return
	}

	l.len--
	if val == l.begin() {
		l.pbeg = l.pbeg.next
		return
	}

	prev := l.begin()
	for cur := l.begin().next; cur != l.end(); cur = cur.next {
		if cur == val {
			prev.next = cur.next
		}
		prev = cur
	}
}

func (l *list) clear() {
	var temp *el
	for cur := l.begin(); cur != l.end(); cur = temp {
		temp = cur.next
		cur.next = nil
		cur = nil
	}
	l.pbeg = nil
	l.len = 0
}

func (l *list) print() {
	for cur := l.begin(); cur != l.end(); cur = cur.next {
		fmt.Printf("%v ", cur.data)
	}
	fmt.Printf("\n")
}

func main() {
	l := list{}

	l.append(10, 5, 5)
	fmt.Print("list after appending: ")
	l.print()
	fmt.Println("len =", l.len)

	l.prepend(1, 2, 3, 4, 5)
	fmt.Print("list after prepending: ")
	l.print()
	fmt.Println("len =", l.len)

	fmt.Print("list after insertion: ")
	for cur := l.begin(); cur != l.end(); cur = cur.next {
		if cur.data == 5 {
			l.insert(7, cur)
		}
	}
	l.print()
	fmt.Println("len =", l.len)

	for cur := l.begin(); cur != l.end(); cur = cur.next {
		if cur.data == 5 {
			l.erase(cur)
		}
	}

	fmt.Print("list after erasing: ")
	l.print()
	fmt.Println("len =", l.len)

	fmt.Print("list after clearing: ")
	l.clear()
	l.print()
	fmt.Println("len =", l.len)
}
