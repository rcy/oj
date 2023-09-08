include .env

start:
	find . -name \*.go -o -name \*.gohtml | entr -r go run .

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

getproddb:
	fly ssh sftp get /data/oj_production.db
	fly ssh sftp get /data/oj_production.db-shm
	fly ssh sftp get /data/oj_production.db-wal

test: export NO_SCHEMA_DUMP=1
test: export SQLITE_DB=:memory:
test:
	go test ./...
