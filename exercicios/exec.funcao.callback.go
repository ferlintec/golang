// - Callback: passe uma função como argumento a outra função.
// - Solução: https://play.golang.org/p/2epLD_Yyap

package main

import "fmt"

func main() {
	funcao(imprimaAlgo)
}

func funcao(f func()) {
	f()
}

func imprimaAlgo() {
	fmt.Println("Ola, tudo bem?")
}
