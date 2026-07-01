package dto

type UpdatePostDto struct {
	Subject  string `json:"subject,omitempty" validate:"omitempty"`
	Description  string `json:"description,omitempty" validate:"omitempty"`
}
