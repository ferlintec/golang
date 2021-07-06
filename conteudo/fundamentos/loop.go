package main

import "fmt"

func main() {

	for x := 0; x < 10; x++ {

		fmt.Println(x)
	}

	//Simulando o while

	y := 0
	for y < 10 {
		fmt.Println("Y =", y)
		y++
	}

	//Simulando for until
	z := 0
	for {
		fmt.Println("Z =", z)
		if z >= 10 {
			fmt.Println("Finalizou for until")
			break
		}
		z++
	}
}
