package main

import "fmt"

func init() {
	fmt.Println("sempre chamada XD")
}

func main() {
	resultado := Soma(1, 3)
	fmt.Println(resultado)

	soma, subtracao, multiplicacao, divisao := Operacoes(1, 2)

	fmt.Println(soma, subtracao, multiplicacao, divisao)
}

// Retorno simples

func Subtracao(valor1, valor2 int) int {
	return valor1 - valor2
}

func Soma(valor1 int, valor2 int) int {
	return valor1 + valor2
}

func Divisao(valor1, valor2 int) int {
	if valor2 == 0 {
		return 0.0
	}
	return valor1 / valor2
}

func Multiplicacao(valor1, valor2 int) int {
	return valor1 * valor2
}

// Retorno composto

func Operacoes(valor1 int, valor2 int) (soma int, subtracao int, multiplicacao int, divisao int) {
	soma = Soma(valor1, valor2)
	subtracao = Subtracao(valor1, valor2)
	multiplicacao = Multiplicacao(valor1, valor2)
	divisao = int(Divisao(valor1, valor2))
	return
}
