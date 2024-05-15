.SILENT:

migrate:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

build:
	docker-compose up  --remove-orphans --build
