package utils

import (
    "encoding/json" 
    "fmt"           
    "log"          
    "net/http"      
    "go-gin-api/db"     
    "go-gin-api/models" 
)

// Função Populate é responsável por popular a tabela 'personagem' com dados da API externa Rick and Morty
func Populate() {
    page := 1 // Inicializa a variável 'page' para controlar a paginação

    // Loop infinito que será interrompido quando não houver mais páginas a processar
    for {
        url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", page)

        // Faz uma requisição GET para a API da Rick and Morty
        resp, err := http.Get(url)
        if err != nil {
            log.Fatalln("Failed to fetch characters", err)
        }
        defer resp.Body.Close() // Garante que o corpo da resposta será fechado após o processamento

        // Define uma estrutura temporária para capturar o resultado da resposta JSON
        var result struct {
            Results []models.Personagem `json:"results"` 
            Info    struct {                            
                Next string `json:"next"`             
            } `json:"info"`
        }

        // Decodifica o corpo da resposta JSON no objeto 'result'
        if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
            log.Fatalln("Failed to decode characters", err) 
        }

        // Itera sobre a lista de personagens retornados
        for _, personagem := range result.Results {
            // Insere cada personagem no banco de dados
            _, err := db.DB.Exec(`INSERT INTO personagem (name, status, species, type, gender, image, url) 
                                  VALUES ($1, $2, $3, $4, $5, $6, $7)`,
                personagem.Name, personagem.Status, personagem.Species, personagem.Type, personagem.Gender, personagem.Image, personagem.Url)
            
            // Se houver erro na inserção de algum personagem, ele é logado mas o loop continua
            if err != nil {
                log.Println("Failed to insert character", personagem.Name, err)
            }
        }

        // Verifica se há uma próxima página. Se o campo 'next' for vazio, o loop é interrompido
        if result.Info.Next == "" {
            break
        }

        page++ // Incrementa a página para buscar a próxima
    }
}
