package app

import (
	"context"
	"culand-internship/internal/faq/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.FAQ, error)
	GetByID(ctx context.Context, id int) (*domain.FAQ, error)
	Create(ctx context.Context, internship *domain.FAQ) (int, error)
	Update(ctx context.Context, internship *domain.FAQ) error
	GetByIDForUpdate(ctx context.Context, id int) (*domain.FAQ, error)
	Delete(ctx context.Context, id int) error
}

type TxRunner interface {
	WithinTx(ctx context.Context, fn func(ctx context.Context, repo Repository) error) error
}
