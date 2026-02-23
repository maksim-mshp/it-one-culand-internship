package app

import (
	"context"
	"culand-internship/internal/internship/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.Internship, error)
	GetByID(ctx context.Context, id int) (*domain.Internship, error)
	Create(ctx context.Context, internship *domain.Internship) (int, error)
	Update(ctx context.Context, internship *domain.Internship) error
	GetByIDForUpdate(ctx context.Context, id int) (*domain.Internship, error)
	GetAllActive(ctx context.Context) ([]*domain.Internship, error)
	UpdateStatus(ctx context.Context, id int, status domain.Status) error
	Delete(ctx context.Context, id int) error
}

type TxRunner interface {
	WithinTx(ctx context.Context, fn func(ctx context.Context, repo Repository) error) error
}
