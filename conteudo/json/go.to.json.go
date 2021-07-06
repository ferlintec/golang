package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	//Quando a estrutura tem letra maiúscula, ele pode ser usado externamente
	type Pessoa struct {
		Nome   string   `json:Nome`
		Idade  int      `json:Idade`
		Filhos []string `json:crianças`
	}

	pe := Pessoa{
		"Adriano",
		42,
		[]string{"Pedro", "Artur", "Helena"},
	}

	//https://golang.org/pkg/encoding/json/#Marshal
	b, err := json.Marshal(pe)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	os.Stdout.Write(b)
}
