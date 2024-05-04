migratedown:
	migrate -database postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable -path db/migrations down

migrateup:
	migrate -database postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable -path db/migrations up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
.PHONY: migratedown migrateup sqlc test