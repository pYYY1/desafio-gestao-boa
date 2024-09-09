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
    // DSN (Data Source Name) contendo as informações de conexão com o PostgreSQL
    dsn := "postgresql://banco_gestao_boa_user:MuyUDYQxJcYsxXU98MJvWw4J5LRZQ14l@dpg-crf5143gbbvc73bvje2g-a.oregon-postgres.render.com/banco_gestao_boa?sslmode=require"
    
    // Conecta ao banco de dados usando sqlx e o DSN fornecido
    DB, err = sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalln("Failed to connect to database:", err)
    }

    // Query para criar a tabela 'personagem' se ela ainda não existir
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS personagem (
        id SERIAL PRIMARY KEY,           
        name VARCHAR(255) NOT NULL,    
        status VARCHAR(50),             
        species VARCHAR(100),          
        type VARCHAR(100),               
        gender VARCHAR(50),              
        image TEXT,                      
        url TEXT                        
    );`

    // Executa a query para criar a tabela
    _, err = DB.Exec(createTableQuery)
    if err != nil {
        log.Fatalln("Failed to create table:", err)
    }
}
