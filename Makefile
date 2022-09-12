dev:
	npx foreman start -s

psql:
	psql ${DATABASE_URL}
