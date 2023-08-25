package main

import (
	log "github.com/sirupsen/logrus"
	//_ "go-starter-gin-gorm/api"
	"go-starter-gin-gorm/internal/app/server"
)

// @title Go-Starter-Gin API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath

// @in header
// @name Authorization
func main() {
	log.Info("Start Go-Starter-Gin-Gorm Service ....")
	//task.CheckHealth()
	err := server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
