package main

import "fmt"

func main() {

	var i int
	var j = i //j será do tipo int

	fmt.Printf("Tipo %T Valor %v\n", j, j) //int

	var cp = 0.32 + 3i

	fmt.Printf("Tipo %T Valor %v\n", cp, cp) //complex128

}
