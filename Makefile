include .env

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	bin/golangci-lint run ./... --config .golangci.pipeline.yaml

install-proto-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

install-gateway-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2

install-validate-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1

install-swagger-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-proto-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

install-mock-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest

generate:
	mkdir -p pkg/swagger
	make generate-user-api
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 --proto_path vendor.protogen \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/user_v1/user.proto

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 --proto_path vendor.protogen \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
    --plugin=protoc-gen-validate=bin/protoc-gen-validate \
	api/auth_v1/auth.proto

generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 --proto_path vendor.protogen \
	--go_out=pkg/access_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/access_v1 --validate_opt=paths=source_relative \
    --plugin=protoc-gen-validate=bin/protoc-gen-validate \
	api/access_v1/access.proto

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-goose-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	bin/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

.PHONY: test
test:
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/GalichAnton/auth/internal/services/...,github.com/GalichAnton/auth/internal/api/...

.PHONY: test-coverage
test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/GalichAnton/auth/internal/services/...,github.com/github.com/GalichAnton/auth/internal/api/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore

vendor-proto:
			@if [ ! -d vendor.protogen/validate ]; then \
				mkdir -p vendor.protogen/validate &&\
				git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
				mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
				rm -rf vendor.protogen/protoc-gen-validate ;\
			fi
			@if [ ! -d vendor.protogen/google ]; then \
				git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
				mkdir -p  vendor.protogen/google/ &&\
				mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
				rm -rf vendor.protogen/googleapis ;\
			fi
			@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
				mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
				git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
				mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
				rm -rf vendor.protogen/openapiv2 ;\
			fi

grpc-load-test:
	ghz \
		--proto api/user_v1/user.proto \
		--import-paths=./vendor.protogen \
		--call UserV1.Get \
		--data '{"id": 1}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50051

grpc-error-load-test:
	ghz \
		--proto api/user_v1/user.proto \
		--import-paths=./vendor.protogen \
		--call UserV1.Get \
		--data '{"id": 0}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50051