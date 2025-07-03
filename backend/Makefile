run: build
	@ echo "starting server..."
	./bin/server

build:
	go build -o ./bin/server ./cmd/api/main.go

export DB_TEST_PATH = $(shell pwd)/test.db

test: migrate_test
	@ go test ./pkg/repositories/sqlite/student/ -v

migrate_test:
	sqlite3 test.db < ./db/migrations/delete_tables.down.sql
	sqlite3 test.db < ./db/migrations/create_tables.up.sql

migrate_prod:
	sqlite3 database.db < ./db/migrations/delete_tables.down.sql
	sqlite3 database.db < ./db/migrations/create_tables.up.sql
