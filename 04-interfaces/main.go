package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return c.radius * c.radius * math.Pi
}

type circulo struct {
	radius float64
}

// Estas funcoes sao utilizadas antes de criarmos interface

func ExibirAreaRetangulo(retangulo retangulo) {
	area := retangulo.largura * retangulo.altura
	fmt.Println(area)
}

func ExibirAreaCirculo(circulo circulo) {
	area := math.Pi * circulo.radius * circulo.radius
	fmt.Println(area)
}

// funcao que criamos apos configurar as interfaces

func ExibirGeometria(g geometria) {
	fmt.Println(g.area())
}

func main() {
	fmt.Println("Inicializando")
	retangulo := retangulo{
		largura: 3,
		altura:  2,
	}
	circulo := circulo{
		radius: 3,
	}
	ExibirAreaRetangulo(retangulo)
	ExibirAreaCirculo(circulo)
	fmt.Println("=============")
	ExibirGeometria(retangulo)
	ExibirGeometria(circulo)

}
