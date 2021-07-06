package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10) //1 << (1 * 10)
	MB
	GB
	TB
)

func main() {

	printIotasConst()

	x := 10
	y := x >> 2

	fmt.Printf("X %b valor %v\n", x, x)
	fmt.Printf("Y %b valor %v\n", y, y)

	x = 24
	y = x << 2

	fmt.Printf("X %b valor %v hexadecimal %#X\n", x, x, x)
	fmt.Printf("Y %b valor %v\n", y, y)
}

func printIotasConst() {

	fmt.Printf("binary\t\t\t\t\t\t\tdecimal\n")
	fmt.Printf("KB: %b\t\t\t\t\t%d\n", KB, KB)
	fmt.Printf("MB: %b\t\t\t\t%d\n", MB, MB)
	fmt.Printf("GB: %b\t\t\t%d\n", GB, GB)
	fmt.Printf("TB: %b\t%d\n", TB, TB)
}
