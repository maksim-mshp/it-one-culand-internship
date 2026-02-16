package postgres

import (
	"context"
	"culand-internship/internal/internship/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Internship, error) {
	const query = `
        SELECT i.id,
               i.title,
               i.label,
               i.description,
               COALESCE(s.skills_arr, '{}') AS skills,
               COALESCE(g.goals_arr, '{}') AS goals
        FROM culand_internship.internships AS i
                 LEFT JOIN (
                      SELECT internship_id, ARRAY_AGG(skill_name ORDER BY id) AS skills_arr
                      FROM culand_internship.internship_skills
                      GROUP BY internship_id
                 ) AS s ON i.id = s.internship_id
                 LEFT JOIN (
                      SELECT internship_id, ARRAY_AGG(goal_name ORDER BY id) AS goals_arr
                      FROM culand_internship.internship_goals
                      GROUP BY internship_id
                 ) AS g ON i.id = g.internship_id
        ORDER BY i.id;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	internships, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Internship])
	if err != nil {
		return nil, err
	}

	return internships, nil
}
