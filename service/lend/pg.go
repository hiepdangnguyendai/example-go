package lend

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
func (s *pgService) Create(_ context.Context, p *domain.Lend) error {
	book := domain.Book{Model: domain.Model{ID: p.BookID}}
	user := domain.User{Model: domain.Model{ID: p.UserID}}
	if err := s.db.Find(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrBookIsNotExisted
		}
		return err
	}
	if err := s.db.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrUserIsNotExisted
		}
		return err
	}
	var lends []domain.Lend
	s.db.Find(&lends)
	for _, v := range lends {
		if v.BookID == p.BookID {
			return ErrBookNotAvailable
		}
	}
	lend := domain.Lend{
		UserID: p.UserID,
		BookID: p.BookID,
		To:     p.To,
	}
	return s.db.Create(&lend).Error
}
func (s *pgService) Update(_ context.Context, p *domain.Lend) (*domain.Lend, error) {
	book := domain.Book{Model: domain.Model{ID: p.BookID}}
	user := domain.User{Model: domain.Model{ID: p.UserID}}
	if err := s.db.Find(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrBookIsNotExisted
		}
		return nil, err
	}
	if err := s.db.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserIsNotExisted
		}
		return nil, err
	}
	var lends []domain.Lend
	s.db.Find(&lends)
	for _, v := range lends {
		if v.BookID == p.BookID {
			return nil, ErrBookNotAvailable
		}
	}
	old := domain.Lend{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	old.BookID = p.BookID
	old.UserID = p.UserID
	old.From = old.From
	old.To = old.To
	return &old, s.db.Save(&old).Error
}
func (s *pgService) Find(_ context.Context, p *domain.Lend) (*domain.Lend, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res, nil
}
func (s *pgService) FindAll(_ context.Context) ([]domain.Lend, error) {
	res := []domain.Lend{}
	return res, s.db.Find(&res).Error
}
func (s *pgService) Delete(_ context.Context, p *domain.Lend) error {
	old := domain.Lend{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
