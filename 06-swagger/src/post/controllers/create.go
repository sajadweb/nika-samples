package controllers

import (
	"nikaapp/src/post/dto"
	res "nikaapp/src/post/response"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

// Create a Post godoc
//
// @Summary Create a Post
// @Description Create a Post
// @Tags Post
// @Accept json
// @Produce json
// @Param request body dto.CreatePostDto true "Post Data"
// @Success 201 {object} res.CreateResponse
// @Failure 400 {object} res.Error "Error code"
// @Failure 422 {object} res.Error "Validation error"
// @Router /posts [post]
func (c *PostController) CreateHandler(ctx *gin.Context) {
	var input dto.CreatePostDto
	if !validator.BindAndValidate(ctx, &input) {
		return
	}
	result, err := c.service.Create(ctx, &input)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Create(
		ctx,
		res.CreateResponse{
			Success: true,
			Message: "create post successfull",
			Data:    result,
		},
	)
}
