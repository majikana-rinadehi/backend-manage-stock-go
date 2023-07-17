MIGRATION_DIR=./db/migrations

run:
	swag init && go mod tidy && go run .

create-migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq create_schema

migrateup:
	go run ./db/main.go up

migratedown:
	go run ./db/main.go down

build:
	go build -o bin/backend-manage-stock-go -v .

build-migration:
	go build -o bin/migration -v ./db

gen-test-handler:
	gotests -template_dir="util/test/templates/handler" -w -all "$(ARG)"

gen-test-handler-method:
	gotests -only=$(ARG) -template_dir="util/test/templates/handler" -w "$(ARG)"