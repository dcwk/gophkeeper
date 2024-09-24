LOCAL_BIN:=$(CURDIR)/bin

protoc: ## Generate package from proto files
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/gophkeeper/gophkeeper.proto

migrate: ## Migrate database to last version
	goose -dir ./migrations/sql postgres "postgres://postgres:123456@localhost:5432/gophkeeper" up