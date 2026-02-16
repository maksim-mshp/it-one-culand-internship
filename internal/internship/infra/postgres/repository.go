package postgres

import (
	"context"
	"culand-internship/internal/internship/domain"
	"errors"
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

func (r *Repository) GetByID(ctx context.Context, id int) (*domain.Internship, error) {
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
        WHERE i.id = $1;
	`

	var internship domain.Internship

	row := r.db.QueryRow(ctx, query, id)
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
	const query = `
		WITH
		new_internship AS (
			INSERT INTO culand_internship.internships (title, label, description)
				VALUES ($1, $2, $3)
				RETURNING id
		),
		insert_skills AS (
			INSERT INTO culand_internship.internship_skills (internship_id, skill_name)
				SELECT new_internship.id, UNNEST($4::TEXT[])
				FROM new_internship
		),
		insert_goals AS (
			INSERT INTO culand_internship.internship_goals (internship_id, goal_name)
				SELECT new_internship.id, UNNEST($5::TEXT[])
				FROM new_internship
		)
		SELECT id FROM new_internship
	`

	var id int
	row := r.db.QueryRow(ctx, query,
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
