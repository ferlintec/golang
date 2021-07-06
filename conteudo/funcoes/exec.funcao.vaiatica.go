// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.
// - Solução: https://play.golang.org/p/Tgv3wwxKV-
package main

import "fmt"

func soma1(slice ...int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func soma2(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func main() {

	s := []int{1, 2, 3, 4, 5}

	fmt.Println("Soma de", s, "é", soma1(s...))
	fmt.Println("Soma de", s, "é", soma2(s))

}
