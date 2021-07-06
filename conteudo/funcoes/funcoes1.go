package main

import "fmt"

func main() {

	soma, quant, s := soma(1, 3, 5, 7)

	fmt.Println("Resultado:", soma, quant, s)
}

func soma(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {

		soma += v
	}
	return soma, len(x), "CQD"
}
