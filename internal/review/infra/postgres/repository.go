package postgres

import (
	"context"
	"culand-internship/internal/core/postgres"
	"culand-internship/internal/review/domain"
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

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Review, error) {
	rows, err := r.db.Query(ctx, GetAllQuery)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*domain.Review, error) {
		var (
			id          int
			textPtr     *string
			authorPtr   *string
			positionPtr *string
		)

		err = row.Scan(&id, &textPtr, &authorPtr, &positionPtr)
		if err != nil {
			return nil, err
		}

		text := utils.ReconstituteVO(textPtr, domain.ReconstituteText)
		author := utils.ReconstituteVO(authorPtr, domain.ReconstituteAuthor)
		position := utils.ReconstituteVO(positionPtr, domain.ReconstitutePosition)
		review := domain.ReconstituteReview(id, text, author, position)

		return &review, nil
	})
}

func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Review, error) {
	var (
		newID       int
		textPtr     *string
		authorPtr   *string
		positionPtr *string
	)

	row := r.db.QueryRow(ctx, GetByIDQuery, id)
	err := row.Scan(&newID, &textPtr, &authorPtr, &positionPtr)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError{ID: id}
		}
		return nil, err
	}

	text := utils.ReconstituteVO(textPtr, domain.ReconstituteText)
	author := utils.ReconstituteVO(authorPtr, domain.ReconstituteAuthor)
	position := utils.ReconstituteVO(positionPtr, domain.ReconstitutePosition)
	review := domain.ReconstituteReview(newID, text, author, position)

	return &review, nil
}

func (r *Repository) Create(ctx context.Context, review *domain.Review) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, CreateQuery,
		utils.EmptyStringToNil(review.Text().Value()),
		utils.EmptyStringToNil(review.Author().Value()),
		utils.EmptyStringToNil(review.Position().Value()),
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, review *domain.Review) error {
	tag, err := r.db.Exec(ctx, UpdateQuery,
		utils.EmptyStringToNil(review.Text().Value()),
		utils.EmptyStringToNil(review.Author().Value()),
		utils.EmptyStringToNil(review.Position().Value()),
		review.ID(),
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.NotFoundError{ID: review.ID()}
	}
	return nil
}

func (r *Repository) GetByIDForUpdate(ctx context.Context, id int) (*domain.Review, error) {
	var (
		newID       int
		textPtr     *string
		authorPtr   *string
		positionPtr *string
	)

	row := r.db.QueryRow(ctx, GetByIDForUpdateQuery, id)
	err := row.Scan(&newID, &textPtr, &authorPtr, &positionPtr)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError{ID: id}
		}
		return nil, err
	}

	text := utils.ReconstituteVO(textPtr, domain.ReconstituteText)
	author := utils.ReconstituteVO(authorPtr, domain.ReconstituteAuthor)
	position := utils.ReconstituteVO(positionPtr, domain.ReconstitutePosition)
	review := domain.ReconstituteReview(newID, text, author, position)

	return &review, nil
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
