package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

// func (receiver) identifier (parameters) (returns) { code }
func (p pessoa) facaAlgo() {
	fmt.Println("Essa funçao existe somente para o tipo pessoa.")
}
func main() {

}
