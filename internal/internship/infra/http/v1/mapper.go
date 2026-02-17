package v1

import "culand-internship/internal/internship/domain"

func MapInternship(i *domain.Internship) InternshipDto {
	return InternshipDto{
		ID:          i.ID,
		Title:       i.Title,
		Label:       i.Label,
		Description: i.Description,
		Skills:      i.Skills,
		Goals:       i.Goals,
	}
}
