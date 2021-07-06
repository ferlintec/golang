package main

import "fmt"

func main() {

	x := retornaUmaFuncao()

	y := x(3)

	fmt.Println(y)

}

func retornaUmaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
