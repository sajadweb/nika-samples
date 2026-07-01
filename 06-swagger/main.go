package main

import (
	_ "nikaapp/docs"
	"nikaapp/src"

	"github.com/sajadweb/nika"
	"github.com/sajadweb/nika/common/config"
	"github.com/sajadweb/nika/common/validator"
	"github.com/sajadweb/nika/common/swagger"
)
// @title Nika API
// @version 0.1
// @description My Nika API
// @host localhost:3007
// @BasePath /
func main() {
	app := nika.NewApp()

	// Validation
	validator.Setup(app)
 
	swagger.Setup(app,&swagger.Config{Path:"/swagger/*any"})
	//envPath := ".env"
	cfg := config.Setup(app, ".env") 
	 
	rootModule := src.NewAppModule()
	app.LoadModule(rootModule)

	port := cfg.Get("PORT","3007")
	app.Listen(":" + port)
}
