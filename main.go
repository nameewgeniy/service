package main

import (
	"github.com/gin-gonic/gin"
	"service/db"
	"service/domain/controller"
	"service/domain/repository"
)

func main() {

	conn := db.GetConnect()
	defer db.CloseConnect(conn)

	// Configuration repositories
	configRep := repository.ConfigRepository{
		DB: conn,
	}

	// Init Repositories
	repository.InitConfig(&configRep)

	// Initial Server
	configCont := controller.ConfigController{
		Server: gin.Default(),
	}

	controller.InitConfig(&configCont)

	configCont.Server.Run()
}
