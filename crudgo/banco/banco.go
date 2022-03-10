package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Pra tirar o warning
)

//ABRINDO A CONEXÃO COM O BANCO DE DADOS
func Conectar() (*sql.DB, error) {
	stringConexao := "root@tcp(localhost)/dbcrud?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
