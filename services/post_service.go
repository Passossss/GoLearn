package services

import (
	"PostGo/models"
	"context"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) GetAll(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	if err := s.db.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetByID(ctx context.Context, id uint) (*models.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) Create(ctx context.Context, post *models.Post) error {
	return s.db.WithContext(ctx).Create(post).Error
}

func (s *PostService) Update(ctx context.Context, id uint, updated *models.Post) error {
	var post models.Post
	if err := s.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return err
	}
	post.Title = updated.Title
	post.Content = updated.Content
	return s.db.WithContext(ctx).Save(&post).Error
}

func (s *PostService) Delete(ctx context.Context, id uint) error {
	if err := s.db.WithContext(ctx).Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
