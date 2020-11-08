package controller

import (
	"github.com/gin-gonic/gin"
	"service/domain/controller/idea_controller"
)

type ConfigController struct {
	Server *gin.Engine
}

func InitRoutes(cf *ConfigController)  {

	// Список всех идей
	cf.Server.GET("/ideas", idea_controller.GetIdeas)

}


