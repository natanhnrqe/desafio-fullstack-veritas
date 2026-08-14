package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "host=localhost port=5433 user=kanban password=kanban123 dbname=kanban sslmode=disable"
	}

	var err error

	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Banco não respondeu:", err)
	}

	log.Println("Banco de dados conectado!")
	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		status VARCHAR(50) NOT NULL DEFAULT 'TODO',
		create_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal("Erro ao criar tabela: ", err)
	}

	log.Println("Tabela criada com sucesso!!!")
}
