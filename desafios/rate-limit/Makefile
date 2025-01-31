api::
	@go build -o bin/app cmd/main.go
	./bin/app

redis::
	@docker-compose -f docker-compose.yml up -d --remove-orphans

start::redis api