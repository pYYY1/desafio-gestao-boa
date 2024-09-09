package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "go-gin-api/db"
    "go-gin-api/models"
)

// Função para criar um novo personagem (rota: POST /personagens)
func CreateCharacter(c *gin.Context) {
    var personagem models.Personagem
    // Tenta vincular o JSON recebido ao struct 'Personagem'
    if err := c.ShouldBindJSON(&personagem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Query SQL para inserir o novo personagem no banco de dados e retornar o ID gerado
    query := `INSERT INTO personagem (name, status, species, type, gender, image, url) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
    err := db.DB.QueryRow(query, personagem.Name, personagem.Status, personagem.Species, personagem.Type, personagem.Gender, personagem.Image, personagem.Url).Scan(&personagem.ID)

    // Verifica se houve erro ao inserir no banco
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert character"})
        return
    }

    // Retorna o ID do personagem criado
    c.JSON(http.StatusOK, gin.H{"id": personagem.ID})
}

// Função para atualizar um personagem existente (rota: PUT /personagens/:personagem_id)
func UpdateCharacter(c *gin.Context) {
    id := c.Param("personagem_id") // Obtém o ID do personagem a ser atualizado a partir da URL

    // Converte o ID de string para inteiro
    idInt, err := strconv.Atoi(id)
    if err != nil {
        // Retorna um erro caso o ID seja inválido
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var personagem models.Personagem

    // Tenta vincular o JSON recebido ao struct 'Personagem'
    if err := c.ShouldBindJSON(&personagem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Query SQL para atualizar os campos do personagem, mantendo os valores existentes caso o campo esteja vazio
    query := `UPDATE personagem SET 
              name = COALESCE(NULLIF($1, ''), name), 
              status = COALESCE(NULLIF($2, ''), status), 
              species = COALESCE(NULLIF($3, ''), species), 
              type = COALESCE(NULLIF($4, ''), type), 
              gender = COALESCE(NULLIF($5, ''), gender), 
              image = COALESCE(NULLIF($6, ''), image), 
              url = COALESCE(NULLIF($7, ''), url)
              WHERE id = $8`

    // Executa a query de atualização no banco de dados
    _, err = db.DB.Exec(query, personagem.Name, personagem.Status, personagem.Species, personagem.Type, personagem.Gender, personagem.Image, personagem.Url, idInt)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update character"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Character updated successfully"})
}

// Função para listar personagens (rota: GET /personagens)
func ListCharacter(c *gin.Context) {
    var personagens []models.Personagem

    query := `SELECT * FROM personagem` // Query SQL base para selecionar todos os personagens

    // Verifica se algum filtro de status ou ordenação foi passado na URL
    status := c.Query("status")
    order := c.Query("order")

    if status != "" {
        query += ` WHERE status = $1` // Adiciona filtro de status
    }

    if order == "asc" {
        query += ` ORDER BY name ASC` // Ordenação crescente
    } else if order == "desc" {
        query += ` ORDER BY name DESC` // Ordenação decrescente
    }

    var err error
    // Executa a query com ou sem filtro de status
    if status != "" {
        err = db.DB.Select(&personagens, query, status)
    } else {
        err = db.DB.Select(&personagens, query)
    }

    // Verifica se houve erro ao listar os personagens
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list characters"})
        return
    }
    c.JSON(http.StatusOK, personagens)
}

// Função para buscar um personagem pelo ID (rota: GET /personagens/:personagem_id)
func GetCharacter(c *gin.Context) {
    id := c.Param("personagem_id") // Obtém o ID do personagem a ser buscado
    var personagem models.Personagem

    // Query SQL para buscar um personagem específico pelo ID
    query := `SELECT * FROM personagem WHERE id = $1`
    err := db.DB.Get(&personagem, query, id)

    // Verifica se o personagem foi encontrado
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }

    // Retorna o personagem encontrado
    c.JSON(http.StatusOK, personagem)
}

// Função para deletar um personagem pelo ID (rota: DELETE /personagens/:personagem_id)
func DeleteCharacter(c *gin.Context) {
    id := c.Param("personagem_id") // Obtém o ID do personagem a ser deletado

    // Query SQL para deletar o personagem pelo ID
    query := `DELETE FROM personagem WHERE id = $1`
    result, err := db.DB.Exec(query, id)

    // Verifica se houve erro ao tentar deletar o personagem
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete character"})
        return
    }

    // Verifica se algum personagem foi realmente deletado (baseado no número de linhas afetadas)
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Character deleted successfully"})
}
