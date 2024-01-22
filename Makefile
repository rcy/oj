-include .env

start: clean-temp
	air

clean-temp:
	rm -rf ${TEMP}/go-build*

build:
	go build .

sql:
	sqlite3 ${SQLITE_DB}

version.%:
	echo "update migration_version set version = $*" | sqlite3 ${SQLITE_DB}

deploy:
	flyctl deploy

drop:
	-rm ${SQLITE_DB}{,-shm,-wal}

seed: db/seed.sql
	sqlite3 ${SQLITE_DB} < $<

db/schema-fixed.sql: db/schema.sql
	sed -e 's/\"//g' $< > $@

generate: db/schema-fixed.sql
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

getproddb:
	fly ssh sftp get /data/oj_production.db
	fly ssh sftp get /data/oj_production.db-shm
	fly ssh sftp get /data/oj_production.db-wal

test:
	. ./.env.test && go test ./...
