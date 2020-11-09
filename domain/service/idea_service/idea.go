package idea_service

import (
	"service/domain/dto"
	"service/domain/repository/idea_repository"
)

func GetIdeas(query * dto.ListQuery) []idea_repository.Idea {

	return idea_repository.GetItems(query)
}

func CreateIdea()  {

	idea := idea_repository.Idea{Title: "Ещё одна идея"}
	idea_repository.Create(idea)
}