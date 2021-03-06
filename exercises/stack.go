package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return

	}
	s.data[s.i] = k
	s.i++

}

func (s *stack) pop() {
	if s.i < 1 {
		return
	}
	s.i--
	s.data[s.i] = 0

}

func main() {
	s := new(stack)
	s.push(25)
	fmt.Printf("stack %v\n", s)

	s.push(8)
	s.push(83)
	s.push(82)
	s.push(81)
	fmt.Printf("stack %v\n", s)
	s.pop()

	fmt.Printf("stack %v\n", s)
}
