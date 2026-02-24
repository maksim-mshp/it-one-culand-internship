package app

type CreateCommand struct {
	Text     *string
	Author   *string
	Position *string
}

type UpdateCommand struct {
	ID       int
	Text     *string
	Author   *string
	Position *string
}

type DeleteCommand struct {
	ID int
}
