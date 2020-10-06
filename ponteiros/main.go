package main

import "fmt"

//Imovel teste
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara", 28000}

	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)
}
