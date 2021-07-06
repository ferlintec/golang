package main

// int int8 int32 int64

// uint uint8 uint32 uint64

// string

// float64 float32

// bool true / false

// complex64 complex128   a + bi

// byte, rune

import "fmt"

type hotdog int

var b hotdog = 12

//Tipo iota (sequencial)
const (
	i1 = iota
	i2 = iota
	i3 = iota
	c1 = 10.0
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

func funcTipos() {

	var altura = 1.34

	var aberto = true

	fmt.Printf("Tipo: %T Valor: %v\n", altura, altura)

	fmt.Printf("Tipo: %T Valor: %v\n", aberto, aberto)

	fmt.Printf("Tipo: %T Valor: %v\n", b, b)

	//inteiro := b #não funciona assim
	inteiro := int(b)

	fmt.Printf("Tipo: %T Valor: %v\n", inteiro, inteiro)

	z := "ß"
	fmt.Printf("Tipo %T Valor %v\n", z, z)

	zb := []byte(z)
	fmt.Printf("Tipo %T Valor %v\n", zb, zb)

}
