package DB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Connect abre conexão com o banco de dados
func Connect() (*sql.DB, error) {
	con := "root:''@/api_go?charset=utf-8"
	db, err := sql.Open("mysql", con)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
