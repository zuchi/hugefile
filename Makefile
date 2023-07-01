test:
	@go test -cover ./...

build-containers:
	@docker-compose build --no-cache

start:
	@docker-compose up -d

stop:
	@docker-compose down