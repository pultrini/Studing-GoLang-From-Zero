package test

import (
	"structures/structures"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := structures.LinkedList{}

	list.InsertAtEnd(10)
	list.Insert(10)

	if list.Length != 2 {
		t.Errorf("Expected 2, recived %d", list.Length)
	}

	if !list.Search(10) {
		t.Error("Value 10 not found")
	}
}
