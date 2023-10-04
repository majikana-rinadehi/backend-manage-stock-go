MIGRATION_DIR=./db/migrations

run:
	swag init && go mod tidy && air -c .air.toml

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
	gotests -template_dir="util/test/templates/handler" -w -all "$(ARG)"

# 全パッケージのテスト実施、カバレッジファイルの出力
do-test: 
	go test ./... -coverpkg=./... -coverprofile=coverage.out > test_report.txt

# カバレッジファイルの表示
coverage-test:
	go tool cover -html=coverage.out -o coverage.html && open coverage.html

test: do-test coverage-test