package v1

type InternshipDto struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
	Goals       []string `json:"goals"`
} // @name InternshipDto

type InternshipRequestDto struct {
	Title       *string   `json:"title"`
	Label       *string   `json:"label"`
	Description *string   `json:"description"`
	Skills      *[]string `json:"skills"`
	Goals       *[]string `json:"goals"`
} // @name InternshipRequestDto
