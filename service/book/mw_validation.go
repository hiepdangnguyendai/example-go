package book

import (
	"context"

	"github.com/hiepdangnguyendai/example-go/domain"
)

type validationMiddleware struct {
	Service
}

func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}
func (mw validationMiddleware) Create(ctx context.Context, book *domain.Book) (err error) {
	switch {
	case len(book.Name) < 5:
		return ErrNameShort
	case book.Name == "":
		return ErrNameEmpty
	}
	return mw.Service.Create(ctx, book)
}

func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Book, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	return mw.Service.Find(ctx, book)
}
func (mw validationMiddleware) Update(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	switch {
	case len(book.Name) < 5:
		return nil, ErrNameShort
	case book.Name == "":
		return nil, ErrNameEmpty
	}
	return mw.Service.Update(ctx, book)
}
func (mw validationMiddleware) Delete(ctx context.Context, book *domain.Book) error {
	return mw.Service.Delete(ctx, book)
}
