package category

import (
	"context"
	"errors"

	"github.com/hiepdangnguyendai/example-go/domain"
	"github.com/jinzhu/gorm"
)

type pgService struct {
	db *gorm.DB
}

func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}
func ValidateNameUnique(s *pgService, name string) error {
	category := domain.Category{}
	if err := s.db.Where("Name = ?", name).Find(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
	}
	return nil
}
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	//return s.db.Create(p).Error
	if err := ValidateNameUnique(s, p.Name); err != nil {
		return s.db.Create(p).Error
	}
	return errors.New("Name has already exited")
}
func (s *pgService) Update(_ context.Context, p *domain.Category) (*domain.Category, error) {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	old.Name = p.Name
	return &old, s.db.Save(&old).Error
}
func (s *pgService) Find(_ context.Context, p *domain.Category) (*domain.Category, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res, nil
}
func (s *pgService) FindAll(_ context.Context) ([]domain.Category, error) {
	res := []domain.Category{}
	return res, s.db.Find(&res).Error
}
func DeleteAllBookBelongCategory(s *pgService, id domain.UUID) error {
	if err := s.db.Where("category_id = ?", id).Delete(&domain.Book{}); err != nil {
		return nil
	}
	return errors.New("Id not existed")
}
func (s *pgService) Delete(_ context.Context, p *domain.Category) error {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	defer DeleteAllBookBelongCategory(s, p.ID)
	return s.db.Delete(old).Error
}
