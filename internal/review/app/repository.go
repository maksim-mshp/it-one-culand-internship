package app

import (
	"context"
	"culand-internship/internal/review/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.Review, error)
	GetByID(ctx context.Context, id int) (*domain.Review, error)
	Create(ctx context.Context, review *domain.Review) (int, error)
	Update(ctx context.Context, review *domain.Review) error
	GetByIDForUpdate(ctx context.Context, id int) (*domain.Review, error)
	Delete(ctx context.Context, id int) error
}

type TxRunner interface {
	WithinTx(ctx context.Context, fn func(ctx context.Context, repo Repository) error) error
}
