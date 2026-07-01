package dto

type FindOnePostDto struct {
	ID string `uri:"id" example:"1" validate:"required"`
}
