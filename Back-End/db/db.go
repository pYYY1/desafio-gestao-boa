package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB // Variável global para armazenar a conexão com o banco de dados

// Função para verificar se a tabela 'personagem' está vazia
func IsTableEmpty() (bool, error) {
	var count int
	// Executa uma query que conta o número de linhas na tabela 'personagem'
	err := DB.QueryRow("SELECT COUNT(*) FROM personagem").Scan(&count)
	if err != nil {
		return false, err
	}
	// Retorna 'true' se a contagem for 0, ou seja, a tabela está vazia
	return count == 0, nil
}

// Função para inicializar a conexão com o banco de dados e criar a tabela, se necessário
func Init() {
	var err error
	dsn := "postgresql://gestao_boa_user:yFtaNqZU0PYR21sHC8CgVrH3LgEAR1FU@dpg-crflcuijhlqs738nno7g-a.oregon-postgres.render.com/gestao_boa?sslmode=require"
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	// Query para criar a tabela 'personagem' se ela ainda não existir
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS personagem (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('Alive', 'Dead', 'unknown')) NOT NULL,
    species VARCHAR(255) NOT NULL,
    type VARCHAR(255),
    gender VARCHAR(50) CHECK (gender IN ('Female', 'Male', 'Genderless', 'unknown')) NOT NULL,
    image VARCHAR(500) NOT NULL,
    url VARCHAR(500) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW()               
    );`

	// Executa a query para criar a tabela
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalln("Failed to create table:", err)
	}
}
