package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetConnect() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), "host=localhost user=go_user password=go_password dbname=go_test_db port=9190")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func CloseConnect(conn *pgx.Conn) {
	conn.Close(context.Background())
}
