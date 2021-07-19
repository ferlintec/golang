package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Pessoa struct {
		Nome  string
		Idade int
	}

	p1 := Pessoa{"Paula", 42}
	sliceBytes, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sliceBytes))
}
