package main

import (
	"encoding/json"
	"fmt"

	"github.com/grassiricardo/structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 28
	casa.SetValor(70000)

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("Valor é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)

	fmt.Println("Casa em Json: ", string(objJSON))
}
