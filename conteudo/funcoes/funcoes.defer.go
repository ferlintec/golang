package main

import "fmt"

func main() {

	//defer deixa pra executar a sentença depois q acabar todo
	//o bloco, ou 1 instrução antes do return
	defer fmt.Println("Imprime 1")
	fmt.Println("Imprime 2")
	fmt.Println("Imprime 3")
	fmt.Println("Imprime 4")
}

// Imprime 2
// Imprime 3
// Imprime 4
// Imprime 1
