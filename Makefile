# Makefile
.PHONY: all

docker-up:
	docker compose -f docker-compose.yml up -d --build --force-recreate

docker-down:
	docker compose -f docker-compose.yml down -v

test:
	docker compose -f docker-compose.yml run test gotestsum -f short-verbose $(args)
