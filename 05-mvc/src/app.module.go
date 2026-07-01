package src

import ( 
	"github.com/sajadweb/nika"
)

type AppModule struct{}

func NewAppModule() *AppModule {
	return &AppModule{}
}

func (m *AppModule) Controllers() []interface{} {
	return []interface{}{
		NewAppController,
	}
}

func (m *AppModule) Providers() []interface{} {
	return []interface{}{
		NewAppService,
	}
}

func (m *AppModule) Imports() []nika.Module {
	return []nika.Module{
		//Import Module
	}
}
