package services

import (
	"PostGo/models"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) GetAll() ([]models.Post, error) {
	var posts []models.Post
	if err := s.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) Create(post *models.Post) error {
	return s.db.Create(post).Error
}

func (s *PostService) Update(id uint, updated *models.Post) error {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		return err
	}
	post.Title = updated.Title
	post.Content = updated.Content
	return s.db.Save(&post).Error
}

func (s *PostService) Delete(id uint) error {
	if err := s.db.Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
