package main

import (
	"fmt"
	"math"
)

type Camisetas struct {
	Tamanho string
	Valor   float64
}

func main() {
	camisas := []Camisetas{
		{Tamanho: "S", Valor: 10},
		{Tamanho: "XXL", Valor: 20},
		{Tamanho: "S", Valor: 4},
		{Tamanho: "XXXXXXXXL", Valor: 5},
		{Tamanho: "XXXXXXXXL", Valor: 25},
	}

	mediaPorTamanho, err := calcularMediaPorTamanho(camisas)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	for tamanho, media := range mediaPorTamanho {
		fmt.Printf("Tamanho: %s, Média: %.2f\n", tamanho, media)
	}
}

/* Função que calcula a média dos preços das camisetas por tamanho */
func calcularMediaPorTamanho(camisas []Camisetas) (map[string]float64, error) {
	mediaPorTamanho := make(map[string]float64)
	countPorTamanho := make(map[string]int)

	for _, camisa := range camisas {
		tamanho := camisa.Tamanho
		valor := camisa.Valor

		mediaPorTamanho[tamanho] += valor
		countPorTamanho[tamanho]++
	}

	for tamanho := range mediaPorTamanho {
		mediaPorTamanho[tamanho] /= float64(countPorTamanho[tamanho])
	}

	return mediaPorTamanho, nil
}
