include .env

start:
	find . -name \*.go -o -name \*.html | entr -r go run .

build:
	go build .

sql:
	sqlite3 ${SQLITE_DB}

deploy:
	flyctl deploy
