// - Crie um struct "pessoa"
// - Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
//     - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
//     - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
//     - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
//     - → x.f is shorthand for (*x).f." ←
//     - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
// - Crie um valor do tipo pessoa;
// - Use a função mudeMe e demonstre o resultado.
// - Solução: https://play.golang.org/p/qiYp9leJcn

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func mudeMeuNome(p *pessoa, nome string) {
	p.nome = nome
	//OU
	// (*p).nome = nome
}

func mudeMeuNome2(p pessoa, nome string) {
	p.nome = nome
}

func main() {

	p := pessoa{"Adriano", 42}

	//Passado um ponteiro (endereço) do objeto
	mudeMeuNome(&p, "Adriano Ferlin")

	fmt.Println(p.nome)

	//Passado uma cópia da referência
	mudeMeuNome2(p, "Jose da Silva")

	fmt.Println(p.nome)
}
