/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: https://play.golang.org/p/nD3TW8VQmH
*/
package main

import "fmt"

func main() {

	x := map[string][]string{
		"Adriano": []string{"futebol", "series", "trilha"},
		"Paula":   []string{"handebol", "series", "trilha", "leitura"},
		"Pedro":   nil,
	}

	x["Pedro"] = []string{"guitarra", "leitura"}

	for nome, valor := range x {
		fmt.Println(nome)
		for indice, v1 := range valor {
			fmt.Println("\t", indice, v1)
		}
	}

	fmt.Println("######################")
	//Remover do mapa
	delete(x, "Pedro")

	for nome, valor := range x {
		fmt.Println(nome)
		for indice, v1 := range valor {
			fmt.Println("\t", indice, v1)
		}
	}
}
