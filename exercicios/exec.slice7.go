/*

- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: https://play.golang.org/p/Gh81-d5tMi
*/
package main

import "fmt"

func main() {

	x := [][]string{
		[]string{"gato", "miau", "coçar"},
		[]string{"cachorro", "auau", "pular"},
		[]string{"galinha", "piu", "ciscar"},
	}

	for _, valor := range x {
		fmt.Println(valor[0])
		for _, v1 := range valor {
			fmt.Println("\t", v1)
		}
	}
}
