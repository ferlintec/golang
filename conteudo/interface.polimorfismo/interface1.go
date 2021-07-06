package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type dentista struct {
	pessoa
	atendimentos int
	salario      float32
}

type arquiteto struct {
	pessoa
	tipoconstrucao string
}

func (x dentista) funcTeste() {
	fmt.Println("Meu nome é", x.nome, " e fiz ", x.atendimentos, " atendimentos.")
}

func (x arquiteto) funcTeste() {
	fmt.Println("Meu nome é", x.nome, " e trabalho com ", x.tipoconstrucao)
}

type gente interface {
	funcTeste()
}

func serhumano(g gente) {
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho: ", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu trabalho com ", g.(arquiteto).tipoconstrucao)

	}
	g.funcTeste()
}

func main() {
	jose := dentista{
		pessoa:       pessoa{"Jose", 34},
		atendimentos: 453,
		salario:      23562.23,
	}

	pedro := arquiteto{
		pessoa:         pessoa{"Pedro", 42},
		tipoconstrucao: "prédios",
	}

	jose.funcTeste()
	pedro.funcTeste()

	fmt.Println(" ")
	//Chamar função que invoca método através da interface.
	serhumano(jose)
	fmt.Println(" ")
	serhumano(pedro)

}
