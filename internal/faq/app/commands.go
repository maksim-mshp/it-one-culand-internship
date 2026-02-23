package app

type CreateCommand struct {
	Question *string
	Answer   *string
}

type UpdateCommand struct {
	ID       int
	Question *string
	Answer   *string
}

type DeleteCommand struct {
	ID int
}
