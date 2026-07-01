package src

import (
	"nikaapp/src/post"

	"github.com/sajadweb/nika"
)

type AppModule struct{}

func NewAppModule() *AppModule {
	return &AppModule{}
}

func (m *AppModule) Controllers() []interface{} {
	return []interface{}{ 
	}
}

func (m *AppModule) Providers() []interface{} {
	return []interface{}{}
}

func (m *AppModule) Imports() []nika.Module {
	return []nika.Module{
		post.NewPostModule(),
	}
}
