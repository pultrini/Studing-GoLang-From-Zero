package test

import (
	"structures/structures"
	"testing"
)

func TestStack(t *testing.T) {
	s := structures.LinkedStack{}

	s.Push(100)
	s.Push(200)

	val, err := s.Pop()
	if err != nil || val != 200 {
		t.Errorf("Error on POP, expectecd 200 obtained: %d", val)
	}
}
