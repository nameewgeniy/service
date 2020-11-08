package idea_repository

import (
	"context"
	"fmt"
	"os"
	"service/domain/repository"
)

var config *repository.ConfigRepository

func SetConfig(configRep *repository.ConfigRepository)  {
	config = configRep
}

func Save() {
	_, err := config.DB.Exec(context.Background(), "INSERT INTO products (title, data) values ('Тестовый', '{}')")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}