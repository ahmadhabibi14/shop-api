setup:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate:
	migrate create -ext sql -dir database/migration $(state)

migrate-up:
	migrate -path database/migration -database "mysql://habi:habi123@tcp(localhost:3306)/shop" -verbose up

migrate-down:
	migrate -path database/migration -database "mysql://habi:habi123@tcp(localhost:3306)/shop" -verbose down
