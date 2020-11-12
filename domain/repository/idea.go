package repository

import (
	"context"
	"fmt"
	"os"
	"service/domain/dto"
	"service/domain/entity"
)

func GetItemsIdeas(query *dto.ListQuery) []entity.Idea {

	var ideas []entity.Idea
	var idea entity.Idea

	sql := "SELECT id, title, created_at FROM ideas t LIMIT $1"

	rows, _ := config.DB.Query(context.Background(), sql, query.Count)

	for rows.Next() {

		err := rows.Scan(&idea.ID, &idea.Title, &idea.CreatedAt)

		if err != nil {
			fmt.Printf("Error")
		}

		ideas = append(ideas, idea)
	}

	return ideas
}

func CreateIdea(i entity.Idea) {

	sql := "INSERT INTO ideas (title) values ($1)"

	_, err := config.DB.Exec(context.Background(), sql, i.Title)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}