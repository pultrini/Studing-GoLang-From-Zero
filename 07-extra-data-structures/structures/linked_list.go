// Package structures fornece implementações de estruturas de dados básicas,
// como listas ligadas e pilhas.
package structures

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) InsertAtEnd(val int) {
	newNode := &Node{Value: val}
	l.Length++

	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) Insert(val int) {
	newNode := &Node{Value: val}
	if l.Head != nil {
		newNode.Next = l.Head
	}
	l.Head = newNode
	l.Length++
}

func (l *LinkedList) Display() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d", current.Value)
		current = current.Next
	}
	fmt.Println(nil)
}

func (l *LinkedList) Remove(value int) {

	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.Length--
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Search(value int) {
	current := l.Head
	for current != nil {
		if current.Value == value {
			fmt.Printf("valor %d encontrado\n", current.Value)
			current = current.Next
		}
	}
	fmt.Println("Valor nao encontrado")
}

func (l *LinkedList) Size() int {
	fmt.Printf("%d", l.Length)
	return l.Length
}

func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}
