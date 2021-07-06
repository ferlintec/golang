package main

import "fmt"

//Com tipo iota, ele preenche automatico
const (
	_ = 5 + iota
	a
	b
	c
)

func main() {

	fmt.Println(a, b, c)
}
