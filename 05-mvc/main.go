package main

import (
	"nikaapp/src"

	"github.com/sajadweb/nika"
	"github.com/sajadweb/nika/common/config"
	"github.com/sajadweb/nika/common/validator"
	"github.com/sajadweb/nika/common/template"
)

func main() {
	app := nika.NewApp()

	// Validation
	validator.Setup(app)

	//envPath := ".env"
	cfg := config.Setup(app, ".env")

	//template
	template.Setup(app, "src/views/*.html")

	rootModule := src.NewAppModule()
	app.LoadModule(rootModule)

	port := cfg.Get("PORT", "3007")
	app.Listen(":" + port)
}
