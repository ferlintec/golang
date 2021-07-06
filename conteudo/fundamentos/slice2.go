package main

import "fmt"

func main() {

	//make([]T, length, capacity)
	s := make([]int, 3, 6)
	fmt.Println("emp:", s)

	fmt.Println("len:", len(s), "cap:", cap(s))

	s[0], s[1], s[2] = 1, 2, 3
	fmt.Println("Slice 1:", s)

	//Não funciona. Precisa aumentar a capacidade
	//s[3]  = 4

	s = append(s, 4)
	fmt.Println("Slice 2:", s)

	fmt.Println("len:", len(s), "cap:", cap(s))

	s = append(s, 5)
	s = append(s, 6)
	s = append(s, 7)

	//A capacidade é dobrada quando se atinge o limite.
	fmt.Println("len:", len(s), "cap:", cap(s))
	fmt.Println("Slice 3:", s)
}

//é fatiando q se deleta um slice
func removeSlice(slice []string, indice int) []string {
	return append(slice[:indice], slice[indice+1:]...)
}
func removeSliceInt(slice []int, indice int) []int {
	return append(slice[:indice], slice[indice+1:]...)
}
