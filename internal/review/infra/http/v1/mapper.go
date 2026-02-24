package v1

import "culand-internship/internal/review/domain"

func MapReview(r *domain.Review) ReviewDto {
	return ReviewDto{
		ID:       r.ID(),
		Text:     r.Text().Value(),
		Author:   r.Author().Value(),
		Position: r.Position().Value(),
	}
}
