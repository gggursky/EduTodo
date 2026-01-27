swag-doc:
	 go run github.com/swaggo/swag/cmd/swag@latest init --generalInfo ./cmd/app/main.go --output ./internal/swagger/docs

build-db:
	docker compose --progress=plain build --no-cache edu_todo_pg
up-db:
	docker compose up -d edu_todo_pg