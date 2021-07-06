package main

import "fmt"

func main() {

	x := 103
	y := 2

	func(x, y int) {
		fmt.Println(x, "vezes 3 é:", x*3)
		fmt.Println(y)
	}(x, y)

}
