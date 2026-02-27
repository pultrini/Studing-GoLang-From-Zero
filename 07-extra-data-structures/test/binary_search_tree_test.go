package test

import (
	"structures/structures"
	"testing"
)

func TestBST(t *testing.T) {
	tree := structures.BinaryTree{}

	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)

	if !tree.Search(25) {
		t.Error("Valor 25 deveria existir na árvore")
	}

	if tree.Search(10) {
		t.Error("Valor 10 não deveria existir na árvore")
	}
}
