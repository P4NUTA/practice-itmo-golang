#!/bin/bash
DSN="user=postgres password=postgres dbname=postgres sslmode=disable port=5432 host=localhost"
ARGS=$(filter-out $@,$(MAKECMDGOALS))

build:
	go build -o bin/main ./cmd

test:
	go test ./...

run:
	go run cmd/main.go

migrate:
	cd ./migrations && \
	goose postgres $(DSN) up

migrate-status:
	cd ./migrations && \
    goose postgres $(DSN) status

migrate-gen:
	cd ./migrations/ && \
    goose postgres $(DSN) create $(ARGS) sql