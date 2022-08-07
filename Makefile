#!make
include .env
export

psql:
	psql "${DATABASE_DSN}"

run:
	go run cmd/api/main.go

runfrontend:
	cd frontend && npm run dev

buildapp:
	cd frontend && npm run build
	go build -o build/out cmd/api/main.go

migratecreate:
	migrate create -ext sql -dir store/migrations -seq $(name)

migrateforce:
	migrate -database "${COCKROACHDB_DSN}" -path store/migrations -verbose force $(version)

migrateup:
	migrate -database "${COCKROACHDB_DSN}" -path store/migrations -verbose up

migratefresh:
	migrate -database "${COCKROACHDB_DSN}" -path store/migrations -verbose down
	migrate -database "${COCKROACHDB_DSN}" -path store/migrations -verbose up