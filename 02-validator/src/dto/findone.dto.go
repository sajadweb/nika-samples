package dto

type FindOneUriDto struct {
	Name string `uri:"name" example:"Nika"  validate:"required"`
}
type FindOneQueryDto struct {
	Name string `form:"name" validate:"required"`
}
type FindOneDto struct {
	Name string `json:"name" example:"Nika"  validate:"required"`
}