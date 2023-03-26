include .env

start: install up
	npx foreman start -s

bootstrap: install up
	${MAKE} -C backend create migrate

up:
	podman-compose up -d

install:
	npm i
	cd backend && npm i
	cd frontend-react && npm i

psql:
	psql ${DATABASE_URL}
