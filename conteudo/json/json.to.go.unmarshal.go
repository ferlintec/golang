package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Quando a estrutura tem letra maiúscula, ele pode ser usado externamente
	type Pessoa struct {
		Nome  string `json:"Nome"`
		Idade int    `json:"Idade"`
		//Dentro do Json, Crianças equivale a Filhos no Go.
		Filhos []string `json:"Crianças"`
	}

	sb := []byte(`{"Nome":"Jose","Idade":40,"Crianças":["Artur","Helena"]}`)

	var p Pessoa
	err := json.Unmarshal(sb, &p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p.Filhos)
	// os.Stdout.Write(p.Filhos)
}
