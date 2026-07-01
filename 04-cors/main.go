package main

import (
	"fmt"
	"nikaapp/src"

	"github.com/sajadweb/nika"
	"github.com/sajadweb/nika/common/config"
	"github.com/sajadweb/nika/common/cors"
	"github.com/sajadweb/nika/common/validator"
)

func main() {
	app := nika.NewApp()
	if app == nil {
		fmt.Println("Error")
	}
	//corse 
	_, err := cors.Setup(app, cors.Config{
		AllowOrigins: []string{
			"*", 
		},
		AllowMethods: []string{
			"*",
		},
		AllowHeaders: []string{
			"*",
		}, 
		AllowCredentials: true,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// Validation
	validator.Setup(app)

	//envPath := ".env"
	cfg := config.Setup(app, ".env")

	rootModule := src.NewAppModule()
	app.LoadModule(rootModule)

	port := cfg.Get("PORT", "3007")
	// app.Use()
	app.Listen(":" + port)
}
