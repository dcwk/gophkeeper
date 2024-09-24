LOCAL_BIN:=$(CURDIR)/bin

protoc: ## Generate package from proto files
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/gophkeeper/gophkeeper.proto
