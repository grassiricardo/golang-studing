package main

import (
	"fmt"

	"github.com/grassiricardo/variaveis/pacotes/operadora"
	"github.com/grassiricardo/variaveis/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário
var NomeDoUsuario = "Grassi"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo: %d\r\n", prefixo.Capital)
	fmt.Printf("Prefixo: %s\r\n", operadora.NomeOperadora)
}
