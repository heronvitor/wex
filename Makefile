
.PHONY: get-mockery
get-mockery:
	@go install github.com/vektra/mockery/v2@v2.38.0

.PHONY: mocks
mocks:
	@mockery --all --keeptree

.PHONY: get-migrate
get-migrate:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: migrate-up
migrate-up:
	@migrate -path ./db/migrations -database $(DB_URL) up

.PHONY: docker
docker:
	@docker-compose up --build -d

.PHONY: db
db:
	@psql $(DB_URL)
