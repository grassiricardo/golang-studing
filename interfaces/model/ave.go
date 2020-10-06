package model

//Galinha representa uma ave tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pato representa uma ave do tipo pato
type Pato interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja representa um som feito por uma galinha
func (a Ave) Cacareja() string {
	return "Cocorico..."
}

//Quack representa um som feito por um pato
func (a Ave) Quack() string {
	return "Quack, quack"
}
