package v1

type InternshipDto struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Label       string   `json:"label,omitempty"`
	Description string   `json:"description,omitempty"`
	Skills      []string `json:"skills,omitempty"`
	Goals       []string `json:"goals,omitempty"`
} // @name InternshipDto

type InternshipAdminDto struct {
	ID          int      `json:"id"`
	Status      string   `json:"status"`
	Title       string   `json:"title"`
	Label       string   `json:"label,omitempty"`
	Description string   `json:"description,omitempty"`
	Skills      []string `json:"skills,omitempty"`
	Goals       []string `json:"goals,omitempty"`
} // @name InternshipAdminDto

type InternshipRequestDto struct {
	Title       *string   `json:"title"`
	Label       *string   `json:"label"`
	Description *string   `json:"description"`
	Skills      *[]string `json:"skills"`
	Goals       *[]string `json:"goals"`
} // @name InternshipRequestDto

type InternshipStatusRequestDto struct {
	Status *string `json:"status"`
} // @name InternshipStatusRequestDto

type InternshipInternalStatus struct {
	IsActive bool `json:"isActive"`
} // @name InternshipInternalStatus
