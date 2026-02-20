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

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*domain.Internship, error) {
		var (
			id      int
			t, l, d string
			s, g    []string
		)

		err = row.Scan(&id, &t, &l, &d, &s, &g)
		if err != nil {
			return nil, err
		}

		title := domain.ReconstituteTitle(t)
		label := domain.ReconstituteLabel(l)
		desc := domain.ReconstituteDescription(d)
		skills := domain.ReconstituteSkills(s)
		goals := domain.ReconstituteGoals(g)
		info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
		internship := domain.ReconstituteInternship(id, info)

		return &internship, nil
	})
}

func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Internship, error) {
	var (
		newId   int
		t, l, d string
		s, g    []string
	)

	row := r.db.QueryRow(ctx, GetByIDQuery, id)
	err := row.Scan(&newId, &t, &l, &d, &s, &g)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewNotFoundError(id)
		}
		return nil, err
	}

	title := domain.ReconstituteTitle(t)
	label := domain.ReconstituteLabel(l)
	desc := domain.ReconstituteDescription(d)
	skills := domain.ReconstituteSkills(s)
	goals := domain.ReconstituteGoals(g)
	info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
	internship := domain.ReconstituteInternship(id, info)

	return &internship, nil
}

func (r *Repository) Create(ctx context.Context, internship *domain.Internship) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, CreateQuery,
		internship.Info().Title().Value(),
		internship.Info().Label().Value(),
		internship.Info().Description().Value(),
		internship.Info().Skills().Value(),
		internship.Info().Goals().Value(),
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, internship *domain.Internship) error {
	tag, err := r.db.Exec(ctx, UpdateQuery,
		internship.Info().Title().Value(),
		internship.Info().Label().Value(),
		internship.Info().Description().Value(),
		internship.Info().Skills().Value(),
		internship.Info().Goals().Value(),
		internship.ID(),
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.NewNotFoundError(internship.ID())
	}
	return nil
}

func (r *Repository) GetByIDForUpdate(ctx context.Context, id int) (*domain.Internship, error) {
	var (
		newId   int
		t, l, d string
		s, g    []string
	)

	row := r.db.QueryRow(ctx, GetByIDForUpdateQuery, id)
	err := row.Scan(&newId, &t, &l, &d, &s, &g)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewNotFoundError(id)
		}
		return nil, err
	}

	title := domain.ReconstituteTitle(t)
	label := domain.ReconstituteLabel(l)
	desc := domain.ReconstituteDescription(d)
	skills := domain.ReconstituteSkills(s)
	goals := domain.ReconstituteGoals(g)
	info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
	internship := domain.ReconstituteInternship(id, info)

	return &internship, nil
}
