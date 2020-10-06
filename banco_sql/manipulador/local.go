package manipulador

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/grassiricardo/banco_sql/model"
	"github.com/grassiricardo/banco_sql/repo"
)

//Local e o manipulador da rota /local
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}

	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Nao foi enviado um numero valido", http.StatusBadRequest)
		fmt.Println("[Ola] Erro na execucao", err.Error())
		return
	}

	sql := "Select country, city from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Nao foi possivel pesquisar este numero", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao", err.Error())
		return
	}

	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Nao foi fazer o binding dos dados do banco na estrutura", http.StatusInternalServerError)
			fmt.Println("[Ola] Erro na execucao", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da pagina", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao", err.Error())
	}
}
