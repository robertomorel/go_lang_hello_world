package main

import "fmt"

// Métodos
// São baseados em estruturas - Structs
type Carro struct {
	Nome string
}

/*
	Atachando (c Carro) ao método andar -> Bind
*/
func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	result1 := soma(10, 20)
	fmt.Println(result1)

	result2, str := soma2(10, 20)
	fmt.Println(result2, str)

	result3 := soma3(10, 20)
	fmt.Println(result3)

	result4 := soma4(10, 20, 10, 40, 60, 1)
	fmt.Println(result4)

	// Função anônima
	// result5 que é uma função, que retorna uma função, que retorna um inteiro
	result5 := func(x ...int) func() int {
		res := 0
		for _, v := range x {
			res += v
		}
		return func() int { return res * res }
	}
	fmt.Println(result5(10, 10, 20, 30, 40)) // Assim ele retorna um end. de memória, pois vai retornar a função
	fmt.Println(result5(10, 10, 20, 30, 40)())

	// Método
	carro := Carro{
		Nome: "BMW",
	}

	carro.andar()
}

// Simples
func soma(a int, b int) int {
	return a + b
}

// Retornando dois elementos
func soma2(a int, b int) (int, string) {
	return a + b, "somou!"
}

// Resultado nomeado
func soma3(a int, b int) (result int) {
	result = a + b
	return
}

// Função variática
func soma4(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}
