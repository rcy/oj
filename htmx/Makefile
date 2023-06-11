export PORT=8080
export SQLITE_DB=./oj.db

start:
	find . -name \*.go -o -name \*.html | entr -r go run .

sql:
	sqlite3 ${SQLITE_DB}
