package main

import "fmt"

func main() {

	x := 10

	//Atribui o endereço da variável
	y := &x

	//Endreço da memória da variável
	fmt.Println(&x)

	fmt.Println(y)

	//Dereference
	//A partir do endereço, pegar o que tem dentro dele, ou seja, seu valor.
	fmt.Println(*y)

	//Mostra o conteúdo do endereço da variável x.
	fmt.Println(*&x)

	fmt.Printf("Tipo de x: %T e tipo de y: %T\n", x, y)

	*y++

	fmt.Println(x, y)
}

// 0xc0000b8010
// 0xc0000b8010
// 10
// 10
// Tipo de x: int e tipo de y: *int
