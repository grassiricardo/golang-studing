package main

import "fmt"

var endereco string
var telefone = "99999-9999"

var quantidade int
var comprou bool
var valor float64
var palavras rune

func main() {
	fmt.Printf("endereco: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
}
