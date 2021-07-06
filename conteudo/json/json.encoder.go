package main

import (
	"encoding/json"
	"os"
)

func main() {

	//Quando a estrutura tem letra maiúscula, ele pode ser usado externamente
	type Pessoa struct {
		Nome  string `json:"Nome"`
		Idade int    `json:"Idade"`
		//Dentro do Json, Crianças equivale a Filhos no Go.
		Filhos []string `json:"Crianças"`
	}

	adriano := Pessoa{"Adriano", 42, []string{"Artur", "Helena"}}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(adriano)
}
