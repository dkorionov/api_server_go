# build project
.PHONY: build
build:
	go build -v ./cmd/apiserver

# run all tests
.PHONY: test
test:
	go test -v -race -timeout 30s ./...

# run migrations
.PHONY: go-migrate
go-migrate:
	 docker run  --network host migrator -path=migrations/ -database "postgresql://docker:docker@localhost:5433/restapi_go?sslmode=disable" up
go-migrate-back:
	docker run  --network host migrator -path=migrations/ -database "postgresql://docker:docker@localhost:5433/restapi_go?sslmode=disable" down

.DEFAULT_GOAL := build