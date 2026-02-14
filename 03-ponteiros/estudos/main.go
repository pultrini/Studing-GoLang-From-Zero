package main

import "fmt"

func main() {
	x := 5
	y1 := x
	// Para visualizar que a alocação de memoria e diferente
	// Ate quando estamos trabalhando com funcoes a memoria muda
	fmt.Println(&x, &y1)
	fmt.Println("IMPRIMIR VALORES FUNCTION")
	ImprimirValores(x, y1)
	
	// Se eu for passar a referencia de memoria uso e comercial e para atribuir valores eu uso o asterisco
	y := &x
	*y = 10
	
	fmt.Println("VALORES DE MEMORIA")
	fmt.Println(&x, y)
	ImprimirValores(x, *y)
	
	
}
func ImprimirValores(x, y int) {
	fmt.Println(&x, &y)
}