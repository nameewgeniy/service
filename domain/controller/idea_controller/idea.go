package idea_controller

import (
	"github.com/gin-gonic/gin"
	"service/domain/dto"
	"service/domain/service/idea_service"
)

var GetIdeas = func(c *gin.Context) {

	// Проверки на роли
	var query dto.ListQuery
	err := c.Bind(&query)

	if err != nil {
		return // or continue, etc.
	}

	idea_service.GetIdeas(&query)

	c.JSON(200, gin.H{
		"message": query,
	})

}