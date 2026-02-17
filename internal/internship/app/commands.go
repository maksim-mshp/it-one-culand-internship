package app

type CreateInternshipCommand struct {
	Title       *string
	Label       *string
	Description *string
	Skills      *[]string
	Goals       *[]string
}

type UpdateInternshipCommand struct {
	ID          int
	Title       *string
	Label       *string
	Description *string
	Skills      *[]string
	Goals       *[]string
}
