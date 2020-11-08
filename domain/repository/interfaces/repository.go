package repository

import "service/domain/repository"

type Repository interface {
	SetConfig(repository *repository.ConfigRepository)
}
