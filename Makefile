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
	`go env GOPATH`/bin/sqlc generate

getproddb:
	fly ssh sftp get /data/oj_production.db
	fly ssh sftp get /data/oj_production.db-shm
	fly ssh sftp get /data/oj_production.db-wal

test: export NO_SCHEMA_DUMP=1
test: export SQLITE_DB=:memory:
test: export ROOT_URL=http://localhost:2001
test:
	go test ./...
