package main

import (
	"fmt"
)

// import (
// 	"fmt"
// 	"math"
// )

//Tipo iota (sequencial)
const (
	i1     = iota
	i2     = iota
	i3     = iota
	c1     = 10.0
	c2 int = 10
)

var v1 = c1

func main() {

	fmt.Printf("\nTipo constante c1 %T valor %v\n", c1, c1)
	fmt.Printf("\nTipo var v1 %T valor %v\n", v1, v1)

	//funcTipos()
	iotaFunc()
}

func iotaFunc() {

	fmt.Println(i1, i2, i3)

	fmt.Printf("\nTipo i1 %T valor %v\n", i1, i1)
	fmt.Printf("\nTipo i2 %T valor %v\n", i2, i2)
	fmt.Printf("\nTipo i3 %T valor %v\n", i3, i3)

	var x = 10
	//i1 = x
	fmt.Printf("\nTipo x %T valor %v\n", x, x)
	fmt.Printf("\nTipo i1 após atribuição: %T valor %v\n", i1, i1)

}
func main2() {

	const Pi = 3.14

	fmt.Println(Pi)

	//Pi = 3 //Dá erro, pois não pode alterar valor de constante.

	//const numero = math.Max(3, 5)
	//#erro - não

}
