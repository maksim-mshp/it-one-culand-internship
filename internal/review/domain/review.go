package domain

type Review struct {
	id       int
	text     Text
	author   Author
	position Position
}

func NewReview(t Text, a Author, p Position) (Review, error) {
	return Review{
		text:     t,
		author:   a,
		position: p,
	}, nil
}

func (r *Review) ID() int {
	return r.id
}
func (r *Review) Text() Text {
	return r.text
}
func (r *Review) Author() Author {
	return r.author
}
func (r *Review) Position() Position {
	return r.position
}

func (r *Review) SetID(id int) {
	r.id = id
}

func ReconstituteReview(id int, t Text, a Author, p Position) Review {
	return Review{
		id:       id,
		text:     t,
		author:   a,
		position: p,
	}
}
