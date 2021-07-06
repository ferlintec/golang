package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Soma total de pares:", somaPares(soma, slice...))
}

func soma(s ...int) int {

	total := 0
	for _, v := range s {
		total += v
	}

	return total
}

func somaPares(f func(x ...int) int, y ...int) int {

	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	return f(slice...)
}
