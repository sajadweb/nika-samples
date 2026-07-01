package controllers

import (
	"github.com/gin-gonic/gin"
	"nikaapp/src/post/services"
)

type PostController struct {
	service  *services.PostService

	Create         func(*gin.Context) `route:"POST:/posts"`
	FindOne        func(*gin.Context) `route:"GET:/posts/:id"`
	Find           func(*gin.Context) `route:"GET:/posts"`
	Delete         func(*gin.Context) `route:"DELETE:/posts/:id"`
	Update         func(*gin.Context) `route:"PUT:/posts/:id"`
}

func NewPostController( service *services.PostService ) *PostController {
	c := &PostController{service:  service}
	c.Create = c.CreateHandler 
	c.Find =  c.FindHandler
	c.FindOne = c.FindOneHandler
	c.Update = c.UpdateHandler
	c.Delete = c.DeleteHandler
	return c
}
