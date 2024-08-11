package database

import (
	"gorm.io/gorm"

	"github.com/lucastomic/qranalyzer/domain"
)

type VisitorRepository interface {
	Create(visitor *domain.Visitor) error
	FindAll() ([]domain.Visitor, error)
}

type gormVisitorRepository struct {
	db *gorm.DB
}

func NewGormVisitorRepository(db *gorm.DB) *gormVisitorRepository {
	return &gormVisitorRepository{db}
}

func (r *gormVisitorRepository) Create(visitor *domain.Visitor) error {
	return r.db.Create(visitor).Error
}

func (r *gormVisitorRepository) FindAll() ([]domain.Visitor, error) {
	var visitors []domain.Visitor
	err := r.db.Find(&visitors).Error
	return visitors, err
}
