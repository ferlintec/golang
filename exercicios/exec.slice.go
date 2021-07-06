/*
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
- Solução: https://play.golang.org/p/tpWDzzlO2l
*/
package main

import "fmt"

func main() {

	s := []int{0, 1, 2, 3, 4, 5, 6}

	for _, valor := range s {

		fmt.Printf("Tipo %T Valor: %v\n", valor, valor)
	}

	fmt.Println(s)      //[0 1 2 3 4 5 6]
	fmt.Println(s[3:5]) //[3 4]
	fmt.Println(s[4:])  //[4 5 6]
	fmt.Println(s[:5])  //[0 1 2 3 4]

}
