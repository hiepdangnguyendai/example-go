package book

import (
	"context"

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
func ValidationCategoryExisted(s *pgService, id domain.UUID) bool {
	var category domain.Category
	if err := s.db.Where("id = ?", id).First(&category).Error; err != nil {
		return false
	}
	return true
}
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	if ValidationCategoryExisted(s, p.Category_Id) {
		return s.db.Create(p).Error
	} else {

		return ErrCategoryNotExisted
	}
}
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	old.Name = p.Name
	old.Author = p.Author
	old.Description = p.Description
	return &old, s.db.Save(&old).Error
}
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	res := []domain.Book{}
	return res, s.db.Find(&res).Error
}
