package idea_service

import (
	"service/domain/dto"
	"service/domain/repository/idea_repository"
)

func GetIdeas(query * dto.ListQuery) {
	idea_repository.Save()
}
