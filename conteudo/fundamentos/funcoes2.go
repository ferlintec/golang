package main

import "fmt"

func main() {

	fmt.Println(calcular(2))

	fmt.Println(calcular2(2))
}

func calcular(a int) (int, int) {

	var quadrado = a * a
	var cubo = a * quadrado

	return quadrado, cubo
}

func calcular2(a int) (quadrado int, cubo int) {

	quadrado = a * a
	cubo = a * quadrado

	return quadrado, cubo
}
