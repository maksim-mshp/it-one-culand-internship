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
	const query = `
        SELECT id, title
        FROM culand_internship.internships
        ORDER BY id
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if clsErr := rows.Close(); clsErr != nil && err == nil {
			err = clsErr
		}
	}(rows)

	var list []*domain.Internship
	for rows.Next() {
		item := &domain.Internship{}
		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
