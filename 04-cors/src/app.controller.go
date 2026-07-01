package src

import (
	"nikaapp/src/dto"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

type AppController struct {
	service      *AppService
	GetApp       func(*gin.Context) `route:"GET:/"`
	SayHello     func(*gin.Context) `route:"GET:/:name"`
	SayPutHello  func(*gin.Context) `route:"PUT:/:name"`
	SayPostHello func(*gin.Context) `route:"POST:/:name"`
}

func NewAppController(service *AppService) *AppController {
	c := &AppController{service: service}

	c.SayHello = func(ctx *gin.Context) {
		var findOneDto dto.FindOneDto
		if !validator.BindAndValidateUri(ctx, &findOneDto) {
			return
		}
		message := c.service.GetWelcomeMessage(findOneDto.Name)
		response.Ok(ctx, gin.H{"message": message})
	}
	c.SayPutHello = c.SayHello
	c.SayPostHello = c.SayHello
	c.GetApp = func(ctx *gin.Context) {

		response.Ok(ctx, gin.H{"message": "Hi Nika 0.1.1"})
	}

	return c
}
