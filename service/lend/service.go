package lend

import (
	"context"

	"github.com/hiepdangnguyendai/example-go/domain"
)

type Service interface {
	Create(ctx context.Context, p *domain.Lend) error
	Update(ctx context.Context, p *domain.Lend) (*domain.Lend, error)
	Find(ctx context.Context, p *domain.Lend) (*domain.Lend, error)
	FindAll(ctx context.Context) ([]domain.Lend, error)
	Delete(ctx context.Context, p *domain.Lend) error
}
