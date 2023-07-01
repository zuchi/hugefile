test:
	@go test -cover ./...

build-containers:
	@docker-compose build

start:
	@docker-compose up -d

stop:
	@docker-compose down