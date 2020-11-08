migration.up:
	migrate -path db/migration -database "postgresql://go_user:go_password@localhost:9190/go_test_db?sslmode=disable" -verbose up 1

migration.down:
	migrate down

migration.make:
	migrate create -ext sql -dir db/migration -seq init_schema