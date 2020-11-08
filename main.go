package main

import (
	"github.com/gin-gonic/gin"
	"service/db"
	"service/domain/controller"
	"service/domain/repository"
	"service/domain/repository/idea_repository"
)

func main() {

	conn := db.GetConnect()
	defer db.CloseConnect(conn)

	// Configuration repositories
	configRep := repository.ConfigRepository{
		DB: conn,
	}

	// Init Repositories
	idea_repository.SetConfig(&configRep)

	// Initial Server
	configCont := controller.ConfigController{
		Server: gin.Default(),
	}

	controller.InitRoutes(&configCont)

	configCont.Server.Run()
}
