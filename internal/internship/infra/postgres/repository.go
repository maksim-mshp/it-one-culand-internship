package postgres

import (
	"context"
	"culand-internship/internal/core/postgres"
	"culand-internship/internal/internship/domain"
	"errors"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db postgres.DBTX
}

func NewRepository(db postgres.DBTX) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Internship, error) {
	rows, err := r.db.Query(ctx, GetAllQuery)
	if err != nil {
		return nil, err
	}

	internships, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Internship])
	if err != nil {
		return nil, err
	}

	return internships, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Internship, error) {
	var internship domain.Internship
	row := r.db.QueryRow(ctx, GetByIDQuery, id)
	err := row.Scan(
		&internship.ID,
		&internship.Title,
		&internship.Label,
		&internship.Description,
		&internship.Skills,
		&internship.Goals,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewNotFoundError(id)
		}
		return nil, err
	}

	return &internship, nil
}

func (r *Repository) Create(ctx context.Context, internship *domain.Internship) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, CreateQuery,
		internship.Title,
		internship.Label,
		internship.Description,
		internship.Skills,
		internship.Goals,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, internship *domain.Internship) error {
	_, err := r.db.Exec(ctx, UpdateQuery,
		internship.Title,
		internship.Label,
		internship.Description,
		internship.Skills,
		internship.Goals,
		internship.ID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.NewNotFoundError(internship.ID)
		}
		return err
	}
	return nil
}

func (r *Repository) GetByIDForUpdate(ctx context.Context, id int) (*domain.Internship, error) {
	var internship domain.Internship
	row := r.db.QueryRow(ctx, GetByIDForUpdateQuery, id)
	err := row.Scan(
		&internship.ID,
		&internship.Title,
		&internship.Label,
		&internship.Description,
		&internship.Skills,
		&internship.Goals,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewNotFoundError(id)
		}
		return nil, err
	}

	return &internship, nil
}
