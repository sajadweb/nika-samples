package dto

type CreatePostDto struct {
	Subject  string `json:"subject" validate:"required,min=1,max=255"`
	Description  string `json:"description" validate:"required,min=1,max=255"`
}
