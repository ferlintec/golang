package main

import "fmt"

func main() {

	var idade int = 35
	//idade = 3
	fmt.Println("A idade é", idade)

	var nome, sexo string = "Adriano", "Masculino"
	var estCivil string = "casado"

	fmt.Println("Meu nome é", nome, ", tenho", idade, "anos.")
	fmt.Printf("Estou", estCivil, ". Sexo\n\n", sexo)

	fmt.Println("############################\n")
	var (
		a     int    = 1
		b            = 2 //int não tipado
		char  string = "String tipado"
		char2        = "String NÃO tipado"
	)

	nr, err := fmt.Println(a, b, char, char2)

	fmt.Printf("\nRetorno println, total %v, err %v\n", nr, err)
	declaracaoVariaveis()
}

//ab := "oi" #Funciona somente dentro de code block

//Variável no escopo da classe.
var y = "Bom dia!"

func declaracaoVariaveis() {

	//Declaração e atribuição de tipo
	//Só funciona dentro de code block
	x := 10

	fmt.Printf("\nTipo %T, e valor é %v", y, y)

	fmt.Printf("\nTipo %T, e valor é %v", x, x)

}
