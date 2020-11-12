package service

import (
	"service/domain/dto"
	"service/domain/entity"
	"service/domain/repository"
)

func GetIdeas(query * dto.ListQuery) []entity.Idea {

	return repository.GetItemsIdeas(query)
}

func CreateIdea(idea entity.Idea)  {

	repository.CreateIdea(idea)
}