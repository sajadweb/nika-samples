package src

import (
	"nikaapp/src/dto"

	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
	"github.com/sajadweb/nika/common/validator"
)

type AppController struct {
	service  *AppService
	GetApp func(*gin.Context) `route:"GET:/"`
	SayHelloFromQuery func(*gin.Context) `route:"GET:/hi"`
	SayHello func(*gin.Context) `route:"GET:/hi/:name"`
	SayHelloFromBody func(*gin.Context) `route:"POST:/hi"`
}

func NewAppController(service *AppService) *AppController {
	c := &AppController{service: service}

	c.SayHelloFromQuery = func(ctx *gin.Context) {
		var findOneDto dto.FindOneQueryDto
		if !validator.BindAndValidateQuery(ctx, &findOneDto) {
			return
		}
		message := c.service.GetWelcomeMessage(findOneDto.Name)
		response.Ok(ctx, gin.H{"message": message})
	}

	c.SayHello = func(ctx *gin.Context) {
		var findOneDto dto.FindOneUriDto
		if !validator.BindAndValidateUri(ctx, &findOneDto) {
			return
		}
		message := c.service.GetWelcomeMessage(findOneDto.Name)
		response.Ok(ctx, gin.H{"message": message})
	}
	c.SayHelloFromBody = func(ctx *gin.Context) {
		var findOneDto dto.FindOneDto
		if !validator.BindAndValidate(ctx, &findOneDto) {
			return
		}
		message := c.service.GetWelcomeMessage(findOneDto.Name)
		response.Ok(ctx, gin.H{"message": message})
	}
	c.GetApp = func(ctx *gin.Context) {
		 
		response.Ok(ctx, gin.H{"message": "Hi Nika 0.1.1"})
	}

	return c
}
