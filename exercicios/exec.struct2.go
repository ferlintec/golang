// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
// - Solução: https://play.golang.org/p/GLK11Q1_x8y

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	p1 := pessoa{"Ana", "Silva",
		[]string{"pera", "abacaxi", "limao"}}

	p2 := pessoa{"Carlos", "Almeida",
		[]string{"pera", "abacaxi", "limao"}}

	mapa := map[string]pessoa{}

	mapa[p1.sobrenome] = p1
	mapa[p2.sobrenome] = p2

	fmt.Println(mapa)
	for sobrenome, valor := range mapa {
		fmt.Println(sobrenome)
		fmt.Println("\t", valor.nome, valor.sobrenome)

		for _, sabor := range valor.sabores {
			fmt.Println("\t\t", sabor)
		}
	}

}
