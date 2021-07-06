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

	s = append(s, 52)
	s = append(s, 53, 54, 55)

	fmt.Println(s)

	y := []int{56, 57, 58}

	s = append(s, y...)

	fmt.Println(s)

}
