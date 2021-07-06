package main

import "fmt"

func main() {

	appendSliceProblem()
	appendSliceOk()
}

//não é recomendável usar 2 slices para adicionar e remover elementos
func appendSliceProblem() {
	fmt.Println("### Slice Não OK")

	s1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("s1:", s1)

	s2 := append(s1[:1], s1[2:]...)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// 	### Slice Não OK
	// s1: [1 2 3 4 5 6]
	// s1: [1 3 4 5 6 6]
	// s2: [1 3 4 5 6]
}

//É recomendável usar o mesmo slice para manipucar os dados
func appendSliceOk() {

	fmt.Println("### Slice OK")

	s1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("s1:", s1)

	s1 = append(s1[:1], s1[2:]...)

	fmt.Println("s1:", s1)

	// 	### Slice OK
	// s1: [1 2 3 4 5 6]
	// s1: [1 3 4 5 6]
}
