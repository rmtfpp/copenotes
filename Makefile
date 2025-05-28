.PHONY: run db dbreset

run:
	clear
	go run cmd/server/main.go

db:
	clear
	sqlitebrowser dbs/copenotes.db &

dbreset:
	clear
	rm dbs/copenotes.db
	sqlite3 dbs/copenotes.db "VACUUM;"