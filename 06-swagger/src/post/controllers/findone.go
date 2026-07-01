package controllers

import (
	"nikaapp/src/post/dto"
	res "nikaapp/src/post/response"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

// FindOne a Post godoc
//
// @Summary FindOne a Post
// @Description FindOne a Post
// @Tags Post
// @Accept json
// @Produce json
// @Param request path dto.FindOnePostDto true "Post Id"
// @Success 200 {object} res.CreateResponse
// @Failure 400 {object} res.Error "Error code"
// @Failure 422 {object} res.Error "Validation error"
// @Router /posts/{id} [get]
func (c *PostController) FindOneHandler(ctx *gin.Context) {
	var input dto.FindOnePostDto
	if !validator.BindAndValidateUri(ctx, &input) {
		return
	}
	result, err := c.service.FindOne(ctx, &input)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
 
	response.Ok(
		ctx,
		res.CreateResponse{
			Success: true,
			Message: "find one post",
			Data:    result,
		},
	)
}
