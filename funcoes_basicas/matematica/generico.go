package matematica

//Calculo calulo de funcao
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador a
func Multiplicador(x int, y int) int {
	return x * y
}
