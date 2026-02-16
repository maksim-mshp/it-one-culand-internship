package app

type CreateInternshipCommand struct {
	Title       string
	Label       string
	Description string
	Skills      []string
	Goals       []string
}
