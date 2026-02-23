package app

type CreateCommand struct {
	Title       *string
	Label       *string
	Description *string
	Skills      *[]string
	Goals       *[]string
}

type UpdateCommand struct {
	ID          int
	Title       *string
	Label       *string
	Description *string
	Skills      *[]string
	Goals       *[]string
}

type UpdateStatusCommand struct {
	ID     int
	Status *string
}
