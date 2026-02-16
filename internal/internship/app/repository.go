package app

import (
	"context"
	"culand-internship/internal/internship/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.Internship, error)
	GetByID(ctx context.Context, id int) (*domain.Internship, error)
	Create(ctx context.Context, internship *domain.Internship) (int, error)
}
