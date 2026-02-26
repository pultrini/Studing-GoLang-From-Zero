package main

import "fmt"

func main() {
	slice := []int{6, 1, 2, 4}
	sliceStr := []string{"ä", "b", "c"}
	newStrings := reverse(sliceStr)
	newInts := reverse(slice)
	fmt.Println(newStrings)
	fmt.Println(newInts)
}

type constraintCustom interface {
	int | string
}

// Pode substituir o T no codigo abaixo que faz o mesmo
func reverse[T int | string](slice []T) []T {
	newInts := make([]T, len(slice))
	newIntsLen := len(slice) - 1
	for i := range len(slice) {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}
