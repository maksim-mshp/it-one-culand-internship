package postgres

const (
	GetAllQuery = `
		SELECT id, question, answer
		FROM internship.faq
		ORDER BY id;
	`

	GetByIDQuery = `
		SELECT id, question, answer
		FROM internship.faq
		WHERE id = $1;
	`

	CreateQuery = `
		INSERT INTO internship.faq (question, answer)
		VALUES ($1, $2)
		RETURNING id;
	`

	UpdateQuery = `
		UPDATE internship.faq
		SET question = $1, answer = $2
		WHERE id = $3
		RETURNING id;
	`

	GetByIDForUpdateQuery = `
		SELECT id, question, answer
		FROM internship.faq
		ORDER BY id
		FOR UPDATE;
	`

	DeleteQuery = `
		DELETE FROM internship.faq
		WHERE id = $1;
	`
)
