package main

import "fmt"

type queue []int

func EmptyQueueRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func (s queue) push(v int) queue {
	s = append(s, v)
	return s
}

func (s queue) pop() (queue, int) {
	defer EmptyQueueRecover()
	len := len(s)
	if len <= 0 {
		panic("Queue is Empty")
	}
	return s[1:], s[0]
}

func main() {

	q := make(queue, 0)
	q = q.push(4)
	q = q.push(5)
	q = q.push(6)
	q, ele := q.pop()
	q = q.push(7)

	q, ele2 := q.pop()

	fmt.Println(q, ele, ele2)
}
