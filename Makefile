dev: up
	${MAKE} up
	npx foreman start -s

up:
	docker compose up -d

psql:
	psql ${DATABASE_URL}
