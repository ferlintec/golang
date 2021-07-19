// - O sort que eu quero não existe. Quero fazer o meu.
// - Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
//     - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
// - Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
// - E aí posso fazer do jeito que eu quiser.
// - Exemplo:
//     - struct carros: nome, consumo, potencia
//     - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
//     - tipo ordenarPorPotencia
//     - tipo ordenarPorConsumo
//     - tipo ordenarPorLucroProDonoDoPosto
// - Go Playground: https://play.golang.org/p/KOIhAsE3OK

package main

import (
	"fmt"
	"sort"
)

//Para que um tipo struct possa utilizar a função Sort, é necessário que ele implemente as funções desta interface: Len, Less e Swap.
type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int { return len(x) }

func (x ordenarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}

func (x ordenarPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {

	carros := []carro{
		{"Jeep", 156, 9},
		{"A3", 189, 7},
		{"Corolla", 136, 12},
	}

	fmt.Println(carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println(carros)

}
