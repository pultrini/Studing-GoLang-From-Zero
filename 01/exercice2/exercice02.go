package main

import "fmt"

func main(){
	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	var somaAte5 int
	var somaAte10 int
	
	for i:=0; i < len(lista); i++ {
		if lista[i] <= 5 {
			somaAte5 += lista[i]
		} else {
			somaAte10 += lista[i]
		}
	}
	fmt.Println(somaAte10)
	fmt.Println(somaAte5)
}
