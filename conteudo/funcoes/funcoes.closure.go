package main

import "fmt"

func main() {

	a := f()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := f()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func f() func() int {
	//closure scope
	//Embora o x esteja declarado fora do escopo da função
	// de callback, mesmo assim ele é utilizado.
	x := 0
	return func() int {
		x++
		return x
	}
}
