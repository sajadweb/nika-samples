package dto

type FindOneDto struct {
	Name string `uri:"name" example:"Nika"  validate:"required"`
}