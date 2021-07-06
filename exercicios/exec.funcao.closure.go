// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

package main

import "fmt"

func main() {

	funcaoRetornada := funcao()

	funcaoRetornada()

}

func funcao() func() {

	x := "oi do closure"

	return func() {
		//x é uma variável declarada fora deste escopo, e mesmo assim pode ser usada, por isso é closure
		fmt.Println(x)
	}
}
