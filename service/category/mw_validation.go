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
	if len(category.Name) <= 5 || category.Name == "" {
		return ErrNameEmptyAndShort
	}
	return mw.Service.Create(ctx, category)
}
