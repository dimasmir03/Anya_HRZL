package main

import (
	"bitcoinmonitor/internal/application"
)

// Swagger

//	@title			Coins Monitor Swagger API
//	@version		1.0
//	@description	Swagger API for Coins Monitor.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	saenkodmitriiol@gmail.com

// @BasePath	/api/v1
func main() {
	app := application.NewApplication()
	app.Run()
}
