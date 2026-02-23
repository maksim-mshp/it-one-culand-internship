package v1

type FAQDto struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
} // @name FAQDto

type FAQRequestDto struct {
	Question *string `json:"question"`
	Answer   *string `json:"answer"`
} // @name FAQRequestDto
