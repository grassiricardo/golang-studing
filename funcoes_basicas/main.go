package main

import (
	"fmt"

	"github.com/grassiricardo/funcoes_basicas/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 3)
	soma := matematica.Soma(2, 3)

	fmt.Printf("Resultado %d\r\n", resultado)
	fmt.Printf("Soma %d\r\n", soma)
}
