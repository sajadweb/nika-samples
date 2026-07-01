package src

import "github.com/sajadweb/nika/common/config"

type AppService struct {
	Cfg *config.Config
}

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) GetWelcomeMessage(name string) string {
	return "Hi "+name
}