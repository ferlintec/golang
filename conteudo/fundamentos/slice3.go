package main

import "fmt"

func main() {

	s := [][]int{
		[]int{1, 2, 3, 4},
		[]int{4, 5, 6},
		[]int{7, 8, 9, 10, 11},
	}

	fmt.Println("len:", len(s), "cap:", cap(s))

	fmt.Println(s[1][2])
	fmt.Println(s[2][3])

}

//é fatiando q se deleta um slice
func removeSlice(slice []string, indice int) []string {
	return append(slice[:indice], slice[indice+1:]...)
}
func removeSliceInt(slice []int, indice int) []int {
	return append(slice[:indice], slice[indice+1:]...)
}
