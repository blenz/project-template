shell-frontend:
	@docker compose exec frontend sh

shell-backend:
	@docker compose exec backend sh

build:
	@docker compose build

run:
	@docker compose up