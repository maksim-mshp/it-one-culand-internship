package v1

import "culand-internship/internal/internship/domain"

func MapInternship(i *domain.Internship) InternshipDto {
	return InternshipDto{
		ID:          i.ID(),
		Title:       i.Info().Title().Value(),
		Label:       i.Info().Label().Value(),
		Description: i.Info().Description().Value(),
		Skills:      i.Info().Skills().Value(),
		Goals:       i.Info().Goals().Value(),
	}
}
