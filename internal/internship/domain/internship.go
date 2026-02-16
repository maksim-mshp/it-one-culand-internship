package domain

type Internship struct {
	ID          int
	Title       string
	Label       string
	Description string
	Skills      []string
	Goals       []string
}

func NewInternship(title, label, description string, skills, goals []string) (*Internship, error) {
	if len(title) < 5 {
		return nil, NewTitleTooShortError(len(title), 5)
	}
	if len(title) > 100 {
		return nil, NewTitleTooLongError(len(title), 100)
	}

	return &Internship{
		Title:       title,
		Label:       label,
		Description: description,
		Skills:      skills,
		Goals:       goals,
	}, nil
}
