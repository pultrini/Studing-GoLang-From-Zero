package model

import (
	"errors"
	"time"
)

type ListarCompras struct {
	Mercado    string
	ListaItens []ItensCompras
	Data       time.Time
}
type ItensCompras struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, listaItens []string) (*ListarCompras, error) {

	if mercado == "" {
		return nil, errors.New("mercado obrigatorio")
	}

	if len(listaItens) == 0 {
		return nil, errors.New("Itens obrigatorios")
	}
	
	
	var itens []ItensCompras

	for _, item := range listaItens {
		itens = append(itens, ItensCompras{Nome: item})
	}

	compras := &ListarCompras{
		Mercado:    mercado,
		ListaItens: itens,
		Data:       time.Now(),
	}
	return compras, nil

}
