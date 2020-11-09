package idea_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	result := idea_service.GetIdeas(&query)

	c.JSON(http.StatusOK, result)
}

var CreateIdea = func(c *gin.Context) {

	idea_service.CreateIdea()
}