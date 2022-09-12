start:
	npx foreman start

install:
	npm install

up:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

psql:
	psql ${DATABASE_URL}

create:
	psql ${ROOT_DATABASE_URL} < create.sql

init:
	npx graphile-migrate init

watch:
	npm run watch

commit:
	npx graphile-migrate commit

uncommit:
	npx graphile-migrate uncommit

reset:
	npx graphile-migrate reset --erase
	${MAKE} create

hardreset: down
	rm -rf .gmrc migrations
