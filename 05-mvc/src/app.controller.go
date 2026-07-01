package src

import (
	"nikaapp/src/dto"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/validator"
)

type AppController struct {
	service  *AppService
	GetApp   func(ctx *gin.Context) `route:"GET:/"`
	SayHello func(ctx *gin.Context) `route:"GET:/:name"`
}

func NewAppController(service *AppService) *AppController {
	c := &AppController{service: service}

	c.SayHello = c.SayHelloHandler
	c.GetApp = func(ctx *gin.Context) {
		println("called")
		ctx.HTML(
			200,
			"index.html",
			gin.H{
				"Name": "Nika 0.1.4",
			})
	}
	return c
}

func (c *AppController) SayHelloHandler(ctx *gin.Context) {
	var findOneDto dto.FindOneDto
	if !validator.BindAndValidateUri(ctx, &findOneDto) {
		return
	}
	println("called")
	ctx.HTML(
		200,
		"index.html",
		gin.H{
			"Name": findOneDto.Name,
		},
	)
}
