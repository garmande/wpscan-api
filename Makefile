.PHONY: scan up down build

scan:
	@. .env && docker compose run --rm wpscan --api-token $$WPSCAN_API_TOKEN

up:
	docker compose up --build -d

down:
	docker compose down

build:
	docker compose build
