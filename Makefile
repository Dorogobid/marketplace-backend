COMPOSE ?= docker-compose -f docker-compose.yml
.PHONY:

up-%:
	$(COMPOSE) build $*
	$(COMPOSE) up -d --force-recreate $*
