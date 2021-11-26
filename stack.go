package main

import "fmt"

type stack []int

func EmptyStackRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func (s stack) push(v int) stack {
	s = append(s, v)
	return s
}

func (s stack) pop() (stack, int) {
	defer EmptyStackRecover()
	len := len(s) - 1
	if len < 0 {
		panic("Stack is Empty")
	}
	return s[:len], s[len]
}

func main() {

	st := make(stack, 0)
	st = st.push(4)
	st = st.push(5)
	st = st.push(6)
	st, ele := st.pop()
	st = st.push(7)

	st, ele2 := st.pop()

	fmt.Println(st, ele, ele2)
}
