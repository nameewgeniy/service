package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

var conn *pgx.Conn

type Data struct {
	ID      int
	Doamin  string `json:"domain"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

func main() {

	conn, err := pgx.Connect(context.Background(), "host=localhost user=go_user password=go_password dbname=go_test_db port=9190")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	getDomain()
}

func addDomain() error {
	_, err := conn.Exec(context.Background(), "insert into data(data) values('test')")
	return err
}

func getDomain() Data {

	var result Data

	err := conn.QueryRow(context.Background(), "SELECT * from data").Scan(&result.ID, &result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return result
}
