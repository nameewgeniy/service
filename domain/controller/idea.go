package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/domain/dto"
	"service/domain/entity"
	"service/domain/repository"
	"service/domain/service"
	"strconv"
)

var GetIdeas = func(c *gin.Context) {

	// Проверки на роли
	var query dto.ListQuery
	err := c.Bind(&query)

	if err != nil {
		return // or continue, etc.
	}

	result := repository.GetItemsIdeas(&query)

	c.JSON(http.StatusOK, result)
}

var CreateIdea = func(c *gin.Context) {

	var idea entity.Idea
	err := c.Bind(&idea)

	if err != nil {
		return // or continue, etc.
	}

	service.CreateIdea(idea)
}

var UpdateIdea = func(c *gin.Context) {


	var idea entity.Idea
	err := c.Bind(&idea)

	if err != nil {
		return // or continue, etc.
	}

	idea.ID, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		return // or continue, etc.
	}

	service.UpdateIdea(idea)
}