MAKEFLAGS += --always-make

shell-frontend:
	@docker compose exec frontend sh

shell-backend:
	@docker compose exec backend sh

build:
	@docker compose build

run:
	@docker compose up

infra:
	@docker compose exec infra sh