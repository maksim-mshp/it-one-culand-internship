package postgres

const (
	GetAllQuery = `
		SELECT id, text, author, position
		FROM internship.review
		ORDER BY id;
	`

	GetByIDQuery = `
		SELECT id, text, author, position
		FROM internship.review
		WHERE id = $1;
	`

	CreateQuery = `
		INSERT INTO internship.review (text, author, position)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	UpdateQuery = `
		UPDATE internship.review
		SET text = $1, author = $2, position = $3
		WHERE id = $4
		RETURNING id;
	`

	GetByIDForUpdateQuery = `
		SELECT id, text, author, position
		FROM internship.review
		WHERE id = $1
		FOR UPDATE;
	`

	DeleteQuery = `
		DELETE FROM internship.review
		WHERE id = $1;
	`
)
