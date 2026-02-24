package postgres

import (
	"context"
	"culand-internship/internal/core/postgres"
	"culand-internship/internal/faq/domain"
	"culand-internship/internal/shared/utils"
	"errors"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db postgres.DBTX
}

func NewRepository(db postgres.DBTX) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.FAQ, error) {
	rows, err := r.db.Query(ctx, GetAllQuery)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*domain.FAQ, error) {
		var (
			id   int
			q, a *string
		)

		err = row.Scan(&id, &q, &a)
		if err != nil {
			return nil, err
		}

		question := utils.ReconstituteVO(q, domain.ReconstituteQuestion)
		answer := utils.ReconstituteVO(a, domain.ReconstituteAnswer)
		faq := domain.ReconstituteFAQ(id, question, answer)

		return &faq, nil
	})
}

func (r *Repository) GetByID(ctx context.Context, id int) (*domain.FAQ, error) {
	var (
		newId int
		q, a  *string
	)

	row := r.db.QueryRow(ctx, GetByIDQuery, id)
	err := row.Scan(&newId, &q, &a)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError{ID: id}
		}
		return nil, err
	}

	question := utils.ReconstituteVO(q, domain.ReconstituteQuestion)
	answer := utils.ReconstituteVO(a, domain.ReconstituteAnswer)
	faq := domain.ReconstituteFAQ(id, question, answer)

	return &faq, nil
}

func (r *Repository) Create(ctx context.Context, faq *domain.FAQ) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, CreateQuery,
		utils.EmptyStringToNil(faq.Question().Value()),
		utils.EmptyStringToNil(faq.Answer().Value()),
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, faq *domain.FAQ) error {
	tag, err := r.db.Exec(ctx, UpdateQuery,
		utils.EmptyStringToNil(faq.Question().Value()),
		utils.EmptyStringToNil(faq.Answer().Value()),
		faq.ID(),
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.NotFoundError{ID: faq.ID()}
	}
	return nil
}

func (r *Repository) GetByIDForUpdate(ctx context.Context, id int) (*domain.FAQ, error) {
	var (
		newId int
		q, a  *string
	)

	row := r.db.QueryRow(ctx, GetByIDForUpdateQuery, id)
	err := row.Scan(&newId, &q, &a)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError{ID: id}
		}
		return nil, err
	}

	question := utils.ReconstituteVO(q, domain.ReconstituteQuestion)
	answer := utils.ReconstituteVO(a, domain.ReconstituteAnswer)
	faq := domain.ReconstituteFAQ(id, question, answer)

	return &faq, nil
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
