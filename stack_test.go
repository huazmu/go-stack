package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Init()

	s.Push("h")
	s.Push("e")
	s.Push("l")
	s.Push("l")
	s.Push("o")
	val := s.Traverse()
	for i := 0; i < len(val); i++ {
		fmt.Printf("%v-", val[i])
	}
	fmt.Printf("\n")
	s.Pop()
	s.Pop()
	s.Pop()
	val = s.Traverse()
	for i := 0; i < len(val); i++ {
		fmt.Printf("%v-", val[i])
	}
	fmt.Printf("\n")
}
