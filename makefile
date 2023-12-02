include .env

.PHONY: test

up:
	docker-compose up -d --build
down:
	docker-compose down
test-unit:
	go test -race $(shell go list ./...) -coverprofile=coverage.out
cover:
	go tool cover -html=coverage.out
mocks:
	go install github.com/vektra/mockery/v2@latest
	mockery