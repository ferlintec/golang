// - Sort serve para ordenar slices.
//     - Como faz?
//     - golang.org/pkg/ → sort
//     - godoc.org/sort → examples
//     - Sort altera o valor original!
// - Exemplo: Ints, Strings.
// - Go Playground:
//     - sort.Strings: https://play.golang.org/p/Rs1NVwmg7h
//     - sort.Ints: https://play.golang.org/p/I2_vsHujZa

// → Segue lá: http://twitter.com/ellenkorbes

package main

import (
	"fmt"
	"sort"
)

func main() {

	//Ordenação de strings
	ss := []string{"abril", "janeiro", "fevereiro", "março", "maio", "junho"}
	sort.Strings(ss)
	fmt.Println(ss)

	//Ordenação de ints
	ii := []int{2, 4, 6, 2, 77, 3, 90}
	sort.Ints(ii)
	fmt.Println(ii)
}
