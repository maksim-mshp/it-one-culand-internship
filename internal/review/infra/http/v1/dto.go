package v1

type ReviewDto struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Author   string `json:"author"`
	Position string `json:"position"`
} // @name ReviewDto

type ReviewRequestDto struct {
	Text     *string `json:"text"`
	Author   *string `json:"author"`
	Position *string `json:"position"`
} // @name ReviewRequestDto
