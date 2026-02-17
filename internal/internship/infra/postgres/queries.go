package postgres

const (
	GetAllQuery = `
        SELECT i.id,
               i.title,
               i.label,
               i.description,
               COALESCE(s.skills_arr, '{}') AS skills,
               COALESCE(g.goals_arr, '{}') AS goals
        FROM internship.internships AS i
                 LEFT JOIN (
                      SELECT internship_id, ARRAY_AGG(skill_name ORDER BY id) AS skills_arr
                      FROM internship.internship_skills
                      GROUP BY internship_id
                 ) AS s ON i.id = s.internship_id
                 LEFT JOIN (
                      SELECT internship_id, ARRAY_AGG(goal_name ORDER BY id) AS goals_arr
                      FROM internship.internship_goals
                      GROUP BY internship_id
                 ) AS g ON i.id = g.internship_id
        ORDER BY i.id;
	`

	GetByIDQuery = `
        SELECT i.id,
               i.title,
               i.label,
               i.description,
               COALESCE(s.skills_arr, '{}') AS skills,
               COALESCE(g.goals_arr, '{}') AS goals
        FROM internship.internships AS i
                 LEFT JOIN (
                      SELECT internship_id, ARRAY_AGG(skill_name ORDER BY id) AS skills_arr
                      FROM internship.internship_skills
                      GROUP BY internship_id
                 ) AS s ON i.id = s.internship_id
                 LEFT JOIN (
                      SELECT internship_id, ARRAY_AGG(goal_name ORDER BY id) AS goals_arr
                      FROM internship.internship_goals
                      GROUP BY internship_id
                 ) AS g ON i.id = g.internship_id
        WHERE i.id = $1;
	`

	CreateQuery = `
		WITH
		new_internship AS (
			INSERT INTO internship.internships (title, label, description)
			VALUES ($1, $2, $3)
			RETURNING id
		),
		insert_skills AS (
			INSERT INTO internship.internship_skills (internship_id, skill_name)
			SELECT new_internship.id, UNNEST($4::TEXT[])
			FROM new_internship
		),
		insert_goals AS (
			INSERT INTO internship.internship_goals (internship_id, goal_name)
			SELECT new_internship.id, UNNEST($5::TEXT[])
			FROM new_internship
		)
		SELECT id FROM new_internship
	`

	UpdateQuery = `
		WITH
		updated_internship AS (
			UPDATE internship.internships
			SET title = $1,
				label = $2,
				description = $3
			WHERE id = $6
			RETURNING id
		),
		delete_old_skills AS (
			DELETE FROM internship.internship_skills
			WHERE internship_id IN (SELECT id FROM updated_internship)
		),
		insert_new_skills AS (
			INSERT INTO internship.internship_skills (internship_id, skill_name)
			SELECT (SELECT id FROM updated_internship), UNNEST($4::TEXT[])
			WHERE EXISTS (SELECT 1 FROM updated_internship)
		),
		delete_old_goals AS (
			DELETE FROM internship.internship_goals
			WHERE internship_id IN (SELECT id FROM updated_internship)
		),
		insert_new_goals AS (
			INSERT INTO internship.internship_goals (internship_id, goal_name)
			SELECT (SELECT id FROM updated_internship), UNNEST($5::TEXT[])
			WHERE EXISTS (SELECT 1 FROM updated_internship)
		)
		SELECT id FROM updated_internship;
	`

	GetByIDForUpdateQuery = `
        WITH
		locked AS (
			SELECT id, title, label, description
			FROM internship.internships
			WHERE id = $1
			FOR UPDATE
		)
		SELECT l.id,
			   l.title,
			   l.label,
			   l.description,
			   COALESCE((
				   SELECT ARRAY_AGG(s.skill_name ORDER BY s.id)
				   FROM internship.internship_skills s
				   WHERE s.internship_id = l.id
			   ), '{}') AS skills,
			   COALESCE((
				   SELECT ARRAY_AGG(g.goal_name ORDER BY g.id)
				   FROM internship.internship_goals g
				   WHERE g.internship_id = l.id
			   ), '{}') AS goals
		FROM locked l;
	`
)
