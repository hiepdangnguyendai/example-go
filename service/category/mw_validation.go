package category

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
func (mw validationMiddleware) Create(ctx context.Context, category *domain.Category) (err error) {
	switch {
	case len(category.Name) <= 5:
		return ErrNameEmptyAndShort
	case category.Name == "":
		return ErrNameEmptyAndShort
	}
	return mw.Service.Create(ctx, category)
}

func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Category, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return mw.Service.Find(ctx, category)
}
func (mw validationMiddleware) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	switch {
	case len(category.Name) <= 5:
		return nil, ErrNameIsRequired
	case category.Name == "":
		return nil, ErrNameEmptyAndShort
	}
	return mw.Service.Update(ctx, category)
}
func (mw validationMiddleware) Delete(ctx context.Context, category *domain.Category) error {
	return mw.Service.Delete(ctx, category)
}
