package main

import "fmt"

func main() {

	x := 10

	funcaoExpressao := func(x int) int {

		return x * 3
	}

	fmt.Println(x, "vezes 3 =", funcaoExpressao(x))
}
