// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
// - Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG

package main

import "fmt"

func main() {

	x := struct {
		nome     string
		idade    int
		hobby    []string
		parentes map[string]string
	}{
		nome:  "Adriano",
		idade: 42,
		hobby: []string{"trilha", "seriados"},
		parentes: map[string]string{
			"Paula":  "esposa",
			"Artur":  "filho",
			"Helena": "filha",
		},
	}

	fmt.Println(x.nome)
	fmt.Println(x.idade)
	for _, hob := range x.hobby {
		fmt.Println("\t", hob)
	}
	for nome, p := range x.parentes {
		fmt.Println("\t", nome, p)
	}

}
