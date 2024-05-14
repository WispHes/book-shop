.SILENT:

migrate:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

build:
	docker-compose up  --remove-orphans --build

run:
	docker run --name=restful -e POSTGRES_PASSWORD='12345' -p 5432:5432 -d --rm postgres