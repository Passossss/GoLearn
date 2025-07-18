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

func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := pc.service.GetAll(c)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao buscar posts")
		return
	}
	utils.JSONSuccess(c, "Posts encontrados", posts)
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := pc.service.GetByID(c, uint(id))
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, "Post não encontrado")
		return
	}
	utils.JSONSuccess(c, "Post encontrado", post)
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Dados inválidos")
		return
	}
	if err := pc.service.Create(c, &post); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao criar post")
		return
	}
	utils.JSONSuccess(c, "Post criado com sucesso", post)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Dados inválidos")
		return
	}
	if err := pc.service.Update(c, uint(id), &post); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao atualizar post")
		return
	}
	utils.JSONSuccess(c, "Post atualizado com sucesso", post)
}

func (pc *PostController) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := pc.service.Delete(c, uint(id)); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Erro ao excluir post")
		return
	}
	utils.JSONSuccess(c, "Post excluído com sucesso", nil)
}
