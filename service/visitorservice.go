package service

import (
	"github.com/lucastomic/qranalyzer/database"
	"github.com/lucastomic/qranalyzer/domain"
)

type VisitorService interface {
	Create(visitor *domain.Visitor) error
	FindAll() ([]domain.Visitor, error)
}

type visitorService struct {
	visitorRepository database.VisitorRepository
}

func NewVisitorService(visitorRepository database.VisitorRepository) *visitorService {
	return &visitorService{visitorRepository}
}

func (s *visitorService) Create(visitor *domain.Visitor) error {
	return s.visitorRepository.Create(visitor)
}

func (s *visitorService) FindAll() ([]domain.Visitor, error) {
	return s.visitorRepository.FindAll()
}
