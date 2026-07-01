package post

import (
	"nikaapp/src/post/controllers"
	"nikaapp/src/post/services"
	"github.com/sajadweb/nika"
)

type PostModule struct{}

func NewPostModule() *PostModule {
	return &PostModule{}
}

func (m *PostModule) Controllers() []interface{} {
	return []interface{}{
		controllers.NewPostController,
	}
}

func (m *PostModule) Providers() []interface{} {
	return []interface{}{ 
		services.NewPostService,
	}
}

func (m *PostModule) Imports() []nika.Module {
	return []nika.Module{}
}
