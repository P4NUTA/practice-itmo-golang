#!/bin/bash
DSN="user=postgres password=postgres dbname=postgres sslmode=disable port=5432 host=localhost"
ARGS=$(filter-out $@,$(MAKECMDGOALS))

PHONY: deps
deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

PHONY: build
build:
	go build -o bin/main ./cmd

PHONY: test
test:
	go test ./...

PHONY: run
run:
	go run cmd/main.go

PHONY: migrate
migrate: deps
	goose -dir migrations postgres $(DSN) up

PHONY: migrate-status
migrate-status: deps
	goose -dir migrations postgres $(DSN) status

PHONY: migrate-gen
migrate-gen: deps
	goose -dir migrations postgres $(DSN) create $(ARGS) sql

PHONY: swagger
swagger:
	swag init -g cmd/main.go