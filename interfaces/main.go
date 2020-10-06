package main

import (
	"fmt"

	"github.com/grassiricardo/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da silva"

	QueroUmCacarejo(jojo)
	QueroUmQuack(jojo)
}

//QueroUmCacarejo funcao para um cacarejo
func QueroUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

//QueroUmQuack funcao para um quack
func QueroUmQuack(p model.Pato) {
	fmt.Println(p.Quack())
}
