package articles

import (
	"gorm.io/gorm"
)

type UseCase interface {
	GetAll() ([]*Article, error)
	Get(Id int64) (*Article, error)
}

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetAll() ([]*Article, error) {
	var articles []*Article

	result := s.DB.Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func (s *Service) Get(id int64) (*Article, error) {
	var article *Article

	result := s.DB.First(&article, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
