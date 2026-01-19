package postgres

import (
	"context"
	"culand-internship/internal/internship/domain"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Internship, error) {
	return []*domain.Internship{}, nil
}
