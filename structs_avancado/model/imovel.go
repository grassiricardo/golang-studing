package model

//Imovel teste
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor teste
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna valor
func (i *Imovel) GetValor() int {
	return i.valor
}
