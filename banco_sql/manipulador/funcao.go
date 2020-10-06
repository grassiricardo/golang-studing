package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao é um manipulador de requisição HTTP
func Funcao(w http.ResponseWriter, t *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador de funcao usando pacote")
}
