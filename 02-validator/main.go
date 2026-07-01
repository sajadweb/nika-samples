package main

import (
	"nikaapp/src"

	"github.com/sajadweb/nika"
	"github.com/sajadweb/nika/common/config"
	"github.com/sajadweb/nika/common/validator"
)

func main() {
	app := nika.NewApp()

	// Validation
	validator.Setup(app)
 
	//envPath := ".env"
	cfg := config.Setup(app, ".env") 
	 
	rootModule := src.NewAppModule()
	app.LoadModule(rootModule)

	port := cfg.Get("PORT","3007")
	app.Listen(":" + port)
}
