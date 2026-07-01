package main

import (
	"nikaapp/src"
	"time"

	"github.com/sajadweb/nika"
	"github.com/sajadweb/nika/common/config"
	"github.com/sajadweb/nika/common/ratelimit"
	"github.com/sajadweb/nika/common/validator"
)

func main() {
	app := nika.NewApp()

	// Validation
	validator.Setup(app)
    _, err := ratelimit.Setup(app, ratelimit.Config{
        Requests: 10,
        Window:   time.Minute,
        Driver:   ratelimit.DriverMemory,
    })
    if err != nil {
        panic(err)
    }
	//envPath := ".env"
	cfg := config.Setup(app, ".env") 
	 
	rootModule := src.NewAppModule()
	app.LoadModule(rootModule)

	port := cfg.Get("PORT","3007")
	app.Listen(":" + port)
}
