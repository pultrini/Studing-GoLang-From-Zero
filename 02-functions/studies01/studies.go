package main

import "fmt"

func init() {
	fmt.Println("sempre chamada XD")
}

func main(){
	resultado := Soma(1,3)
	fmt.Println(resultado)
	
	soma, subtracao, multiplicacao, divisao := Operacoes(1, 2)
	fmt.Println(soma, subtracao, multiplicacao, divisao)
}
// Retorno simples
func Soma(valor1 int, valor2 int) int {
	return valor1+valor2
}

// Retorno composto
func Operacoes(valor1 int, valor2 int) (soma int, subtracao int, multiplicacao int, divisao int) {
	soma = valor1+valor2
	subtracao = valor1 - valor2
	multiplicacao = valor1*valor2
	divisao = valor1/valor2
	return
}