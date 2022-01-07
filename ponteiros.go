package main

import "fmt"

func main() {

	// memória-enredeço(10) <------------ a <------------ 10
	a := 10
	fmt.Println(&a) // Print do end. de memória do a
	fmt.Println(a, "está no endereço", &a)

	// Todo endereço de memória fica catalogado nos ponteiros
	// O apontamento que tá mandando a a para a memória será guardado
	var ponteiro *int = &a

	// Reference
	fmt.Println(ponteiro)  // Printa o endereço da memória guardado
	fmt.Println(*ponteiro) // Printa o valor que está sendo apontado
	fmt.Println("O ponteiro está apontando para o valor", *ponteiro, "guardado no endereço", ponteiro)

	*ponteiro = 50
	fmt.Println(*ponteiro) // Joguei 50 no mesmo endereço do a
	// Logo a = 50
	fmt.Println(a)

	b := &a
	*b = 60
	fmt.Println(*b)

	// ---------------------------
	variavel := 10
	pont(&variavel) // Passando por parâmetro um end. de memória
	fmt.Println(variavel)
}

// Recebe um ponteiro inteiro
func pont(a *int) {
	*a = 200 // Mudando o valor que está na memória &variavel
}
