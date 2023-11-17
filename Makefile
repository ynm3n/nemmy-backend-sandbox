include .env

up.nocache:
	docker compose build --no-cache
	docker compose up -d

up:
	docker compose up -d

down:
	docker compose down

psql:
	docker compose exec db psql -h localhost -d ${POSTGRES_USER} -U ${POSTGRES_USER}
