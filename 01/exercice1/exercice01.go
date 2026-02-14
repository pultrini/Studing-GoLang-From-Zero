package main

import "fmt"

func main() {
	list := []int{1, 2}
	var soma int
	for _, value := range list {
		soma += value
	}

	fmt.Println(soma)
}
