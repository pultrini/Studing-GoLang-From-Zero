package structures

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedStack struct {
	Top  *Node
	Size int
}

func (s *LinkedStack) Push(val int) {
	newNode := &Node{Value: val}
	newNode.Next = s.Top
	s.Size++
	s.Top = newNode
}

func (s *LinkedStack) Pop() (int, error) {
	if s.Top == nil {
		return 0, errors.New("stack vazia")
	}
	value := s.Top.Value
	s.Top = s.Top.Next
	s.Size--
	return value, nil
}

func (s *LinkedStack) Display() {
	current := s.Top
	fmt.Print("Topo ->")
	for current != nil {
		fmt.Printf("[%d]", current.Value)
		current = current.Next
	}
	fmt.Println("-> Base")
}
