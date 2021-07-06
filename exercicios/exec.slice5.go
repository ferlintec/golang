package main

import "fmt"

func main() {

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s = append(s[:3], s[5:]...)

	//OU
	// x := s[:3]
	// y := s[5:]
	//s = append(x, y...)

	s = append(s[:3], s[5:]...)

	fmt.Println(s)
	//[0 1 2 5 6 7 8 9 10]
}
