COMPOSE ?= docker-compose -f docker-compose.yml
.PHONY: docs sqlc up

up:
	$(COMPOSE) build marketplace-backend && $(COMPOSE) up -d --force-recreate marketplace-backend

docs:
	swag fmt && swag init --parseDependency

sqlc:
	sqlc generate