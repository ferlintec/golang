package main

import "fmt"

func soma(s ...int) int {

	total := 0
	for _, v := range s {
		total += v
	}

	return total
}

func somaImpares(fsoma func(x ...int) int, s ...int) int {

	var sImpares []int
	for _, v := range s {
		if v%2 != 0 {
			sImpares = append(sImpares, v)
		}
	}
	total := fsoma(sImpares...)

	return total
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Soma dos impares: ", somaImpares(soma, s...))
}
