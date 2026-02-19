package exercicio1

import (
	"exercicios/exercicio1/model"
	"fmt"
	"time"
)

func Exercicio1() {
	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "Arroz")
	nomeDosItens = append(nomeDosItens, "banana")
	nomeDosItens = append(nomeDosItens, "mandioca")

	compras, err := model.NewCompra("Seu ze", time.Now(), nomeDosItens)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*compras)
}
