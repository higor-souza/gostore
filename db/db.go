package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBanoDeDados() *sql.DB {
	conexao := "user=test dbname=pricing_clj password=test host=localhost sslmode=disable port=54321"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
