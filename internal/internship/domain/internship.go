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
	internship := Internship{
		Title:       title,
		Label:       label,
		Description: description,
		Skills:      skills,
		Goals:       goals,
	}
	err := internship.Validate()
	if err != nil {
		return nil, err
	}
	return &internship, nil
}

func (i *Internship) Validate() error {
	if len(i.Title) < 5 {
		return NewTitleTooShortError(len(i.Title), 5)
	}
	if len(i.Title) > 100 {
		return NewTitleTooLongError(len(i.Title), 100)
	}
	return nil
}
