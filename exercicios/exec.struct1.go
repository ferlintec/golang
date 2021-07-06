// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
// - Solução: https://play.golang.org/p/Pyp7vmTJfY

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	p1 := pessoa{"Ana", "Paula",
		[]string{"pera", "abacaxi", "limao"}}

	fmt.Println(p1.nome, p1.sobrenome)
	for _, valor := range p1.sabores {
		fmt.Println("\t", valor)
	}

}
