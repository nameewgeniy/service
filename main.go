package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "host=localhost user=point_user password=point_password dbname=point_db_test port=9882")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int64

	err = conn.QueryRow(context.Background(), "SELECT id from ghost_point_orders").Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id)
}
