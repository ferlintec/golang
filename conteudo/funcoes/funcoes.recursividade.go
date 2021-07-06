package main

import "fmt"

func main() {

	x := 4
	fmt.Println("O fatorial de", x, "é", factorial(x))
	fmt.Println("O fatorial de", x, "é", loops(x))

}

func factorial(x int) int {

	if x == 1 {
		return x
	}

	return x * factorial(x-1)
}

func loops(x int) int {

	total := x

	for x > 2 {
		x -= 1
		total *= x
	}

	return total
}
