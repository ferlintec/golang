package main

import "fmt"

func main() {

	fmt.Println("# Linha 1")
	defer fmt.Println("# Linha 2") //última
	defer fmt.Println("# Linha 3") //penúltima
	f2()
	defer f3()                     //ante penúltima
	defer fmt.Println("# Linha 4") //ante ante penúltima
	fmt.Println("# Linha 5")

}

func f2() {
	fmt.Println("## f2 Linha 1")
	defer fmt.Println("## f2 Linha 2")
	fmt.Println("## f2 Linha 3")
}

func f3() {
	fmt.Println("### f3 Linha 1")
	defer fmt.Println("### f3 Linha 2")
	fmt.Println("### f3 Linha 3")
}

// # Linha 1
// ## f2 Linha 1
// ## f2 Linha 3
// ## f2 Linha 2
// # Linha 5
// # Linha 4
// ### f3 Linha 1
// ### f3 Linha 3
// ### f3 Linha 2
// # Linha 3
// # Linha 2
