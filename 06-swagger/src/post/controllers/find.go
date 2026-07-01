package controllers

import (
	"nikaapp/src/post/dto"
	res "nikaapp/src/post/response"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

// Find all Post godoc
//
// @Summary Find all Post
// @Description Find all Post
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} res.ListResponse
// @Failure 400 {object} res.Error "Error code"
// @Failure 422 {object} res.Error "Validation error"
// @Router /posts [get]
func (c *PostController) FindHandler(ctx *gin.Context) {
	var input dto.ListPostDto
	if !validator.BindAndValidateQuery(ctx, &input) {
		return
	}
	result, err := c.service.Find(ctx, &input)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Ok(
		ctx,
		res.ListResponse{
			Success: true,
			Message: "find one post",
			Data:    result,
		},
	)
}
