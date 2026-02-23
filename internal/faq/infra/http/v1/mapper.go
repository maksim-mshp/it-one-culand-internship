package v1

import "culand-internship/internal/faq/domain"

func MapFAQ(f *domain.FAQ) FAQDto {
	return FAQDto{
		ID:       f.ID(),
		Question: f.Question().Value(),
		Answer:   f.Answer().Value(),
	}
}
