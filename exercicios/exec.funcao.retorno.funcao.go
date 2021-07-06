// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
// - Solução: https://play.golang.org/p/A74rufv6Rs

package main

import "fmt"

func main() {

	funcao := retornaUmaFuncao()
	s := funcao(2)

	fmt.Println(s)
}

func retornaUmaFuncao() func(a int) int {

	return func(a int) int {
		return a * a
	}
}
