package main

import (
	"fmt"
	"net/http"

	"github.com/grassiricardo/banco_sql/repo"
)

func main() {
	err := repo.AbreConexao()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados", err.Error())
		return
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Ola mundo")
	// })

	http.HandleFunc("/local/", manipulador.local)
	// http.HandleFunc("/ola", manipulador.Ola)

	// fmt.Println("Estou escutando, comandante...")
	// http.ListenAndServe(":8181", nil)
}
