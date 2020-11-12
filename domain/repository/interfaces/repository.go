package repository

import "service/domain/repository"

type Repository interface {
	InitConfig(repository *repository.ConfigRepository)
}
