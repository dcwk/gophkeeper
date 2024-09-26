run-server:
	go run ./cmd/server/main.go

protoc: ## Generate package from proto files
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/gophkeeper/gophkeeper.proto

migrate-up: ## Migrate database to one version up
	goose -dir ./migrations/sql postgres "postgres://postgres:123456@localhost:5432/gophkeeper" up

migrate-down: ## Migrate database to one version down
	goose -dir ./migrations/sql postgres "postgres://postgres:123456@localhost:5432/gophkeeper" down