package stack

import (
	"strconv"
)

//stack item
type Stack struct {
	i    int
	data [10]int
}

//method to push an item in the stack
func (s *Stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

//method to pop an item from the stack
func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}

//a formated output for the stack
func (s Stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" +
			strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
