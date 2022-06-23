package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("postgres", config.StringConexaoBanco)
	fmt.Println(config.StringConexaoBanco)
	if erro != nil {
		return nil, erro

	}
	if erro = db.Ping(); erro != nil {

		db.Close()
		return nil, erro

	}

	return db, nil
}
