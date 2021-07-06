package main

import "fmt"

func main() {

	slice := []int{1, 3, 5, 7}

	soma := soma(slice)

	fmt.Println("Resultado:", soma)
}

func soma(x []int) int {
	soma := 0
	for _, v := range x {

		soma += v
	}
	return soma
}
