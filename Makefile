start-postgres:
	docker run -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d --name postgresdb -v simple-bank-vol:/var/lib/postgresql/data postgres:13.3-alpine
createdb:
	docker exec postgresdb createdb --username=root --owner=root simple-bank
dropdb:
	docker exec postgresdb dropdb simple-bank
migrateup:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down

.PHONY: start-postgres createdb dropdb migrateup migratedown