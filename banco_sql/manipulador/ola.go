package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grassiricardo/servidor_web/model"
)

//Ola e o manipulador da rota /ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := ModeloOla.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da pagina", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao", err.Error())
	}
}
