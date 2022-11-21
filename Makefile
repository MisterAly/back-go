postgres: 
	docker run --name pokedb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123mudar -d postgres:12-alpine

createdb:
	docker exec -it pokedb createdb --username=root --owner=root pokedb

dropdb:
	docker exec -it pokedb dropdb pokedb

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:123mudar@localhost:5432/pokedb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123mudar@localhost:5432/pokedb?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
