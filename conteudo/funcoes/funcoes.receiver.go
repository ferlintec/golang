package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

//Seria o equivalente à um método do tipo, como se fosse uma classe.
func (p pessoa) apresentacao() {
	fmt.Println("Eu sou um médoto de Pessoa!")
	fmt.Print("Meu nome é ", p.nome, " e eu tenho ", p.idade, " anos.")
}

func main() {

	jose := pessoa{"Jose", 42}

	jose.apresentacao()

}
