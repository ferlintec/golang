package main

import "fmt"

func main() {

	s := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50,
	}
	fmt.Println(s)
}
