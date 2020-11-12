package controller

import (
	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	Server *gin.Engine
}

func InitConfig(cf *ConfigController)  {

	// Список всех идей
	cf.Server.GET("/ideas", GetIdeas)
	cf.Server.GET("/create", CreateIdea)
}


