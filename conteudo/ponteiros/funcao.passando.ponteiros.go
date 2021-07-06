package main

import "fmt"

func main() {

	x := 0

	fmt.Println(x, &x)

	funcaoValor(x)
	fmt.Println(x, &x)
	ponteiroVariavel(&x)
	fmt.Println(x, &x)

	// 0 0xc00001e0e8
	// 0 0xc00001e0e8
	// 1 0xc00001e0e8

}

//Passa a cópia do valor
func funcaoValor(x int) {
	x++
}

//Passa a cópia do endereço (referência)
func ponteiroVariavel(p *int) {

	//Usa dereference para acessar o valor do ponteiro.
	*p++
}
