package idea_repository

import (
	"context"
	"fmt"
	"os"
	"service/domain/dto"
	"service/domain/repository"
	"time"
)

var config *repository.ConfigRepository

type Idea struct {
	ID int
	Title string
	CreatedAt time.Time
}

func SetConfig(configRep *repository.ConfigRepository)  {
	config = configRep
}

func GetItems(query *dto.ListQuery) []Idea{

	var ideas []Idea
	var idea Idea

	sql := "SELECT id, title, created_at FROM ideas t LIMIT $1"

	rows, _ := config.DB.Query(context.Background(), sql, query.Count)

	for rows.Next() {

		err := rows.Scan(&idea.ID, &idea.Title, &idea.CreatedAt)

		if err != nil {
			fmt.Printf("Error")
		}

		fmt.Println(idea.ID)
		ideas = append(ideas, idea)
	}

	return ideas
}

func Create(i Idea) {

	sql := "INSERT INTO ideas (title) values ($1)"

	_, err := config.DB.Exec(context.Background(), sql, i.Title)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}