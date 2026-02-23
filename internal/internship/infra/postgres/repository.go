package postgres

import (
	"context"
	"culand-internship/internal/core/postgres"
	"culand-internship/internal/internship/domain"
	"culand-internship/internal/shared/utils"
	"errors"
	"github.com/jackc/pgx/v5"
	"strings"
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
			id          int
			t, l, d, st *string
			s, g        *[]string
		)

		err = row.Scan(&id, &t, &l, &d, &st, &s, &g)
		if err != nil {
			return nil, err
		}

		title := utils.ReconstituteVO(t, domain.ReconstituteTitle)
		label := utils.ReconstituteVO(l, domain.ReconstituteLabel)
		desc := utils.ReconstituteVO(d, domain.ReconstituteDescription)
		status := utils.ReconstituteVO(st, domain.ReconstituteStatus)
		skills := utils.ReconstituteVO(s, domain.ReconstituteSkills)
		goals := utils.ReconstituteVO(g, domain.ReconstituteGoals)
		info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
		internship := domain.ReconstituteInternship(id, info, status)

		return &internship, nil
	})
}

func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Internship, error) {
	var (
		newId       int
		t, l, d, st *string
		s, g        *[]string
	)

	row := r.db.QueryRow(ctx, GetByIDQuery, id)
	err := row.Scan(&newId, &t, &l, &d, &st, &s, &g)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError{ID: id}
		}
		return nil, err
	}

	title := utils.ReconstituteVO(t, domain.ReconstituteTitle)
	label := utils.ReconstituteVO(l, domain.ReconstituteLabel)
	desc := utils.ReconstituteVO(d, domain.ReconstituteDescription)
	status := utils.ReconstituteVO(st, domain.ReconstituteStatus)
	skills := utils.ReconstituteVO(s, domain.ReconstituteSkills)
	goals := utils.ReconstituteVO(g, domain.ReconstituteGoals)
	info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
	internship := domain.ReconstituteInternship(id, info, status)

	return &internship, nil
}

func (r *Repository) Create(ctx context.Context, internship *domain.Internship) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, CreateQuery,
		utils.EmptyStringToNil(internship.Info().Title().Value()),
		utils.EmptyStringToNil(internship.Info().Label().Value()),
		utils.EmptyStringToNil(internship.Info().Description().Value()),
		utils.EmptyStringToNil(strings.ToLower(internship.Status().Value())),
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
		utils.EmptyStringToNil(internship.Info().Title().Value()),
		utils.EmptyStringToNil(internship.Info().Label().Value()),
		utils.EmptyStringToNil(internship.Info().Description().Value()),
		internship.Info().Skills().Value(),
		internship.Info().Goals().Value(),
		internship.ID(),
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.NotFoundError{ID: internship.ID()}
	}
	return nil
}

func (r *Repository) GetByIDForUpdate(ctx context.Context, id int) (*domain.Internship, error) {
	var (
		newId       int
		t, l, d, st *string
		s, g        *[]string
	)

	row := r.db.QueryRow(ctx, GetByIDForUpdateQuery, id)
	err := row.Scan(&newId, &t, &l, &d, &st, &s, &g)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError{ID: id}
		}
		return nil, err
	}

	title := utils.ReconstituteVO(t, domain.ReconstituteTitle)
	label := utils.ReconstituteVO(l, domain.ReconstituteLabel)
	desc := utils.ReconstituteVO(d, domain.ReconstituteDescription)
	status := utils.ReconstituteVO(st, domain.ReconstituteStatus)
	skills := utils.ReconstituteVO(s, domain.ReconstituteSkills)
	goals := utils.ReconstituteVO(g, domain.ReconstituteGoals)
	info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
	internship := domain.ReconstituteInternship(id, info, status)

	return &internship, nil
}

func (r *Repository) GetAllActive(ctx context.Context) ([]*domain.Internship, error) {
	rows, err := r.db.Query(ctx, GetAllActiveQuery)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*domain.Internship, error) {
		var (
			id          int
			t, l, d, st *string
			s, g        *[]string
		)

		err = row.Scan(&id, &t, &l, &d, &st, &s, &g)
		if err != nil {
			return nil, err
		}

		title := utils.ReconstituteVO(t, domain.ReconstituteTitle)
		label := utils.ReconstituteVO(l, domain.ReconstituteLabel)
		desc := utils.ReconstituteVO(d, domain.ReconstituteDescription)
		status := utils.ReconstituteVO(st, domain.ReconstituteStatus)
		skills := utils.ReconstituteVO(s, domain.ReconstituteSkills)
		goals := utils.ReconstituteVO(g, domain.ReconstituteGoals)
		info := domain.ReconstituteInternshipInfo(title, label, desc, skills, goals)
		internship := domain.ReconstituteInternship(id, info, status)

		return &internship, nil
	})
}

func (r *Repository) UpdateStatus(ctx context.Context, id int, status domain.Status) error {
	s := strings.ToLower(status.Value())
	tag, err := r.db.Exec(ctx, UpdateStatusQuery, id, s)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.NotFoundError{ID: id}
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx, DeleteQuery, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.NotFoundError{ID: id}
	}
	return nil
}
