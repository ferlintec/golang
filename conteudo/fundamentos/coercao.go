package main

import "fmt"

func main() {

	//tipos1()
	sliceBytes()

}

func tipos1() {

	var a int = 54
	var b float64 = 6.4

	var c int = int(b)

	//c = b  	#Não pode atribuit float a int
	fmt.Println("O valor de C é", c)

	var d float32
	d = float32(a)

	fmt.Println("O valor de D é", d)

	z := "ß"
	fmt.Printf("Tipo %T Valor %v\n", z, z)

	zb := []byte(z)
	fmt.Printf("Tipo %T Valor %v\n", zb, zb)

}

func sliceBytes() {
	s := "Ola mundo!"
	sliceByte := []byte(s)

	fmt.Printf("%v\n%T\n\n", sliceByte, sliceByte)

	//Printa por caracter
	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}
	fmt.Println(" ")

	//Printa por byte. Dá diferença do de cima.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
