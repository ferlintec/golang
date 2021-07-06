package main

import "fmt"

func main() {

	//Maps não são ordenados.
	amigos := map[string]int{
		"jose":  1234,
		"ana":   2323,
		"pedro": 3444,
	}
	fmt.Println("amigos:", amigos)

	fmt.Println(amigos["ana"])

	//Adicionar no map
	amigos["Maria"] = 5678
	fmt.Println("amigos:", amigos)

	//comma ok idiom
	//Este parâmetro retorna true ou false caso a chave seja
	//encontrada no map

	valor, temNoMapa := amigos["ricardo"]

	fmt.Println("Valor:", valor, "Tem no mapa:", temNoMapa)
	//Valor: 0 Tem no mapa: false

	if _, temNoMapa := amigos["ricardo"]; temNoMapa {
		fmt.Println("Tem no mapa")
	} else {
		fmt.Println("Não tem no mapa")
	}

}
