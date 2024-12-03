# make migrate.create name=migration_name
migrate.create:
	goose -dir migrations create ${name} sql
migrate.up:
	goose -dir migrations -allow-missing postgres "host=localhost port=5432 user=admin password=1234 dbname=main sslmode=disable" up
migrate.down:
	goose -dir migrations -allow-missing postgres "host=localhost port=5432 user=admin password=1234 dbname=main sslmode=disable" down
migrate.reset:
	goose -dir migrations -allow-missing postgres "host=localhost port=5432 user=admin password=1234 dbname=main sslmode=disable" down
migrate.status:
	goose -dir migrations -allow-missing postgres "host=localhost port=5432 user=admin password=1234 dbname=main sslmode=disable" status
docker.build:
	docker build --no-cache -t test --build-arg MODE=local --build-arg MODULE=backend .
build:
	go build -race -v cmd/main.go
generate:
	go generate ./...