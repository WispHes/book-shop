.SILENT:

build:
	docker-compose up  --remove-orphans --build

lint:
	golangci-lint run