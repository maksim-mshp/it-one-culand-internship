package domain

type FAQ struct {
	id       int
	question Question
	answer   Answer
}

func NewFAQ(q Question, a Answer) (FAQ, error) {
	return FAQ{
		question: q,
		answer:   a,
	}, nil
}

func (f *FAQ) ID() int {
	return f.id
}
func (f *FAQ) Question() Question {
	return f.question
}
func (f *FAQ) Answer() Answer {
	return f.answer
}

func (f *FAQ) SetID(id int) {
	f.id = id
}

func ReconstituteFAQ(id int, q Question, a Answer) FAQ {
	return FAQ{
		id:       id,
		question: q,
		answer:   a,
	}
}
