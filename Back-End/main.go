package main

import (
    "go-gin-api/db"       
    "go-gin-api/handlers"  
    "go-gin-api/utils"      
    "log"                  
    "github.com/gin-contrib/cors"  
    "github.com/gin-gonic/gin"     
)

func main() {
    // Inicializa a conexão com o banco de dados
    db.Init()

    // Verifica se a tabela "personagem" está vazia
    isEmpty, err := db.IsTableEmpty()
    if err != nil {
        log.Fatalln("Failed to check if table is empty", err)
    }

    // Se a tabela estiver vazia, chama a função "Populate" para preencher com dados iniciais
    if isEmpty {
        utils.Populate()
    }

    // Cria uma instância do router Gin
    r := gin.Default()

    // Configura as políticas de CORS, permitindo que a API seja acessada apenas da origem "http://127.0.0.1:3000"
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:3000"},  
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  
        AllowHeaders:     []string{"Origin", "Content-Type"},  
        ExposeHeaders:    []string{"Content-Length"}, 
        AllowCredentials: true,  
    }))

    // Define as rotas (endpoints) da API
    r.POST("/personagens", handlers.CreateCharacter)          
    r.PUT("/personagens/:personagem_id", handlers.UpdateCharacter) 
    r.GET("/personagens", handlers.ListCharacter)               
    r.GET("/personagens/:personagem_id", handlers.GetCharacter) 
    r.DELETE("/personagens/:personagem_id", handlers.DeleteCharacter) 

    r.Run(":8080")
}
