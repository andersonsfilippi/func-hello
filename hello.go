package main

import "fmt"

type Carro struct {
	Nome string
}

// bind com carro
// método andar()
func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	// retorno normal
	resultado := soma(10, 20)
	fmt.Println(resultado)

	// retorno múltiplo
	resultado, msg := somaM(30, 40)
	fmt.Println(resultado, msg)

	// retorno nomeado
	resultado = somaN(50, 60)
	fmt.Println(resultado)

	// função variádica
	resultado = somaTudo(3, 5, 10, 4, 6, 343)
	fmt.Println(resultado)

	// função dentro de função
	calculo := func(x ...int) func() int {
		res := 0
		// blank para ignorar a chave do array
		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}
	fmt.Println(calculo(54, 54, 54, 54)())

	// chamada para método andar que está em Carro
	carro := Carro{
		Nome: "BMW",
	}
	carro.andar()

}

func soma(a int, b int) int {
	return a + b
}

func somaM(a int, b int) (int, string) {
	return a + b, "somou!"
}

func somaN(a int, b int) (abacaxi int) {
	abacaxi = a + b
	return
}

// função variádica recebe um array conforme o tipo especificado
func somaTudo(x ...int) int {
	resultado := 0
	// blank para ignorar a chave do array
	for _, v := range x {
		resultado += v
	}
	return resultado
}
