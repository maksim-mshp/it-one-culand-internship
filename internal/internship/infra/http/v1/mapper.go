package v1

import "culand-internship/internal/internship/domain"

func MapInternship(i *domain.Internship) InternshipResponse {
	return InternshipResponse{
		ID:    i.ID,
		Title: i.Title,
	}
}
