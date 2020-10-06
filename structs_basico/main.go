package main

import "fmt"

//Imovel Strutura
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("Casa %+v\r\b\n", casa)

	apartamento := Imovel{17, 76, "AP", 76000}
	fmt.Printf("AP %+v\r\b\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		X:     66,
		valor: 45,
	}

	fmt.Printf("Chacara %+v\r\b\n", chacara)

	casa.Nome = "Opa"
	casa.valor = 150
	casa.X = 10
	casa.Y = 20
	fmt.Printf("Casa %+v\r\b\n", casa)

}
