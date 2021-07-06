package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	for indice, valor := range s {
		fmt.Println("No indice", indice, "tem o valor", valor)
	}

	fmt.Println("Entre as posiçoes s[2:4] de", s, "tem", s[2:4])
	//Entre as posiçoes s[2:4] de [a b c d e f] tem [c d]

	//Remove slice posição 2
	s = append(s[:2], s[3:]...)
	fmt.Println("Removido posição 2 do slice:", s)

	fmt.Println("Removido posição 1 do slice:", removeSlice(s, 1))

	s2 := []string{"x", "y", "z"}

	//Adicionar todo array numa slice
	fmt.Println("s:", s)
	s = append(s, s2...)
	fmt.Println("Adicionando slices:", s)

	fmt.Println("Capacity slices:", cap(s), cap(s2))

	sliceLenCapExample()
}

//é fatiando q se deleta um slice
func removeSlice(slice []string, indice int) []string {
	return append(slice[:indice], slice[indice+1:]...)
}
func removeSliceInt(slice []int, indice int) []int {
	return append(slice[:indice], slice[indice+1:]...)
}

func sliceLenCapExample() {

	fmt.Println("### sliceLenCapExample:")
	s := make([]int, 3, 5)
	fmt.Println("Capacity slices:", cap(s), cap(s))
	s = append(s, []int{1, 2, 3})
}
