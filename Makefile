.PHONY: run
run:
		@go run cmd/server/main.go

.PHONY: build
build:
	@templ generate

.PHONY: start_compose
start_compose:
	@docker compose up -d

.PHONY: stop_compose
stop_compose:
	@docker compose stop

.PHONY: start
start:	build run
		@echo "Tryhard App starting..."