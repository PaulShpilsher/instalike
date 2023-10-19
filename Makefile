.PHONY: build
build:
	go build -o bin/ ./cmd/webservice
	cp ./cmd/webservice/.env bin/


.PHONY: postgres
postgres:
	docker run --rm -itd --network host \
		-e POSTGRES_USER=pusr \
		-e POSTGRES_PASSWORD=pusr_secret \
		-e POSTGRES_DB=instalike-data \
		-p 5432:5432 -v /postgres-instalike-data:/var/lib/postgresql/data \
		--name instalike-pg \
		postgres:latest

.PHONY: migration
migration:
	migrate create -ext sql -dir ./database/migrations -seq ${seq}

.PHONY: up
up:
	migrate -path database/migrations -database "postgresql://pusr:pusr_secret@localhost:5432/instalike-data?sslmode=disable" up

.PHONY: down
down:
	migrate -path database/migrations -database "postgresql://pusr:pusr_secret@localhost:5432/instalike-data?sslmode=disable" down 1