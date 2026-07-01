package controllers

import (
	"nikaapp/src/post/dto"
	res "nikaapp/src/post/response"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

// Delete a Post godoc
//
// @Summary Delete a Post
// @Description Delete a Post
// @Tags Post
// @Accept json
// @Produce json
// @Param request path dto.FindOnePostDto true "Post Id"
// @Success 200 {object} res.DeleteResponse
// @Failure 400 {object} res.Error "Error code"
// @Failure 422 {object} res.Error "Validation error"
// @Router /posts/{id} [delete]
func (c *PostController) DeleteHandler(ctx *gin.Context) {
	var input dto.FindOnePostDto
	if !validator.BindAndValidateUri(ctx, &input) {
		return
	}
	err := c.service.Delete(ctx, &input)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}

	response.Ok(
		ctx,
		res.DeleteResponse{
			Success: true,
			Message: "delete one post", 
		},
	)
}
