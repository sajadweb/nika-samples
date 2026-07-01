package controllers

import (
	"nikaapp/src/post/dto"
	res "nikaapp/src/post/response"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

// Update a Post godoc
//
// @Summary Update a Post
// @Description Update a Post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param request body dto.UpdatePostDto true "Post Data"
// @Success 202 {object} res.CreateResponse
// @Failure 400 {object} res.Error "Error code"
// @Failure 422 {object} res.Error "Validation error"
// @Router /posts/{id} [put]
func (c *PostController) UpdateHandler(ctx *gin.Context) {
	var body dto.UpdatePostDto
	var dto dto.FindOnePostDto
	
	if !validator.BindAndValidateUri(ctx, &dto) {
		return
	}
	if !validator.BindAndValidate(ctx, &body) {
		return
	}
	result, err := c.service.Update(ctx, &body, &dto)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Update(
		ctx,
		res.CreateResponse{
			Success: true,
			Message: "update post successfull",
			Data:    result,
		},
	)
}
