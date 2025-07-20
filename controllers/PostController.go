package controllers

import (
	"PostGo/models"
	"PostGo/services"
	"PostGo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	service *services.PostService
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		service: services.NewPostService(db),
	}
}

// GetPosts godoc
// @Summary Lista todos os posts
// @Description Obtém uma lista de todos os posts disponíveis.
// @Tags Posts
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]models.Post} "Lista de posts retornada com sucesso"
// @Failure 500 {object} utils.ErrorResponse "Erro interno do servidor"
// @Router /post [get]
func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := pc.service.GetAll()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao buscar posts")
		return
	}
	utils.JSONSuccess(c, "Posts encontrados", posts)
}

// GetPostByID godoc
// @Summary Obtém um post por ID
// @Description Obtém os detalhes de um post específico pelo seu ID.
// @Tags Posts
// @Produce json
// @Param id path int true "ID do Post"
// @Success 200 {object} utils.SuccessResponse{data=models.Post} "Post encontrado com sucesso"
// @Failure 404 {object} utils.ErrorResponse "Post não encontrado"
// @Failure 500 {object} utils.ErrorResponse "Erro interno do servidor"
// @Router /post/{id} [get]
func (pc *PostController) GetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Melhor tratar o erro
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "ID de post inválido")
		return
	}
	post, err := pc.service.GetByID(uint(id))
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, "Post não encontrado")
		return
	}
	utils.JSONSuccess(c, "Post encontrado", post)
}

// CreatePost godoc
// @Summary Cria um novo post
// @Description Adiciona um novo post ao banco de dados.
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Dados do Post para criar"
// @Success 201 {object} utils.SuccessResponse{data=models.Post} "Post criado com sucesso"
// @Failure 400 {object} utils.ErrorResponse "Dados inválidos"
// @Failure 500 {object} utils.ErrorResponse "Erro interno do servidor"
// @Router /post [post]
func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Dados inválidos")
		return
	}
	if err := pc.service.Create(&post); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao criar post")
		return
	}
	utils.JSONSuccess(c, "Post criado com sucesso", post)
}

// UpdatePost godoc
// @Summary Atualiza um post existente
// @Description Atualiza os dados de um post específico pelo seu ID.
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path int true "ID do Post"
// @Param post body models.Post true "Novos dados do Post"
// @Success 200 {object} utils.SuccessResponse{data=models.Post} "Post atualizado com sucesso"
// @Failure 400 {object} utils.ErrorResponse "Dados inválidos"
// @Failure 500 {object} utils.ErrorResponse "Erro ao atualizar post"
// @Router /post/{id} [put]
func (pc *PostController) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "ID de post inválido")
		return
	}
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Dados inválidos")
		return
	}
	if err := pc.service.Update(uint(id), &post); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao atualizar post")
		return
	}
	utils.JSONSuccess(c, "Post atualizado com sucesso", post)
}

// DeletePost godoc
// @Summary Exclui um post
// @Description Exclui um post específico pelo seu ID.
// @Tags Posts
// @Produce json
// @Param id path int true "ID do Post"
// @Success 200 {object} utils.SuccessResponse "Post excluído com sucesso"
// @Failure 500 {object} utils.ErrorResponse "Erro ao excluir post"
// @Router /post/{id} [delete]
func (pc *PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "ID de post inválido")
		return
	}
	if err := pc.service.Delete(uint(id)); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao excluir post")
		return
	}
	utils.JSONSuccess(c, "Post excluído com sucesso", nil)
}
