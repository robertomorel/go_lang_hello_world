package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

// Método Exibe atachado ao Cliente
func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método: ", c.Nome)
}

// type ClienteInternacional struct {
// 	Nome  string
// 	Email string
// 	Pais  string
// 	CPF   int
// }

type ClienteInternacional struct {
	Cliente
	Pais string `json:"_pais"` // Usando a tag do Go para fazer com que a propriedade apareça como "_pais" no json
}

func main() {
	cliente := Cliente{
		Nome:  "Roberto",
		Email: "roberto@email.com",
		CPF:   123456,
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"John Doe", "j.doe@email.com", 123456}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d", cliente2.Nome, cliente2.Email, cliente2.CPF)
	fmt.Println()

	// cliente3 := ClienteInternacional{
	// 	Nome:  "Jane Doe",
	// 	Email: "jane.doe@email.com",
	// 	Pais:  "EUA",
	// 	CPF:   123456,
	// }
	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Jane Doe",
			Email: "jane.doe@email.com",
			CPF:   123456,
		},
		Pais: "EUA",
	}
	fmt.Printf("Nome: %s. Email: %s. País: %s. CPF: %d", cliente3.Nome, cliente3.Email, cliente3.Pais, cliente3.CPF)
	fmt.Println()

	cliente.Exibe()
	cliente2.Exibe()
	// Herança
	cliente3.Exibe()

	// -------------------------------------- JSON ---------------------------------
	// Marshal pega o que tá no struct e transforma em json
	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(clienteJson)         // Bytes
	fmt.Println(string(clienteJson)) // Transformando bytes em string

	jsonCliente4 := `{"Nome":"Jane Doe2","Email":"jane.doe2@email.com","CPF":12345678,"_pais":"CAN"}`
	cliente4 := ClienteInternacional{}

	// Se deixasse dessa forma, só iria mudar o cliente4 dentro do escopo da função, não de maneira geral
	// json.Unmarshal([]byte(jsonCliente4), cliente4)

	json.Unmarshal([]byte(jsonCliente4), &cliente4) // Para alterar no lugar da memória e não ficar apenas no escopo da função

	fmt.Println(cliente4)
}
