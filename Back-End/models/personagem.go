package models

import "time"

// Estrutura Personagem que representa o modelo de dados de um personagem no sistema
type Personagem struct {
    ID      int       `db:"id" json:"id"`          
    Name    string    `db:"name" json:"name"`       
    Status  string    `db:"status" json:"status"`   
    Species string    `db:"species" json:"species"` 
    Type    string    `db:"type" json:"type"`       
    Gender  string    `db:"gender" json:"gender"`  
    Image   string    `db:"image" json:"image"`     
    Url     string    `db:"url" json:"url"`         
    Created time.Time `db:"created" json:"created"` 
}
