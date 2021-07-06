package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	fumante   bool
}

type profissional struct {
	pessoa
	cargo   string
	salario int
}

func main() {

	pessoa1 := pessoa{
		nome:      "Adriano",
		sobrenome: "Ferlin",
		fumante:   false,
	}

	pessoa2 := pessoa{"Paula", "Borsatto", false}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	fmt.Println(pessoa2.nome)

	gerente := profissional{
		pessoa: pessoa{
			nome:      "Jose",
			sobrenome: "Silva",
			fumante:   true,
		},
		cargo:   "gerente",
		salario: 23000,
	}

	fmt.Println(gerente)
	fmt.Println(gerente.pessoa.nome)

	//Neste caso, como "nome" não é o nome de um campo externo, ele pode ser
	//chamado assim.
	fmt.Println(gerente.nome)
}
