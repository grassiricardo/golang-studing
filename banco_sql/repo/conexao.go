package repo

import (
	/*
		_ "github.com/go-sql-driver/msql" é usado diretamente pela aplicacao
	*/
	_ "github.com/go-sql-driver/msql"
	"github.com/jmoiron/sqlx"
)

//Db armazena a conexao com o banco
var Db *sqlx.DB

//AbreConexao Abre conexao
func AbreConexao() (err error) {
	err = nil
	Db, err = sqlx.Open("", "")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
