.PHONY: cert
cert:
	openssl genrsa -out ./keys/rsa 4096
	openssl rsa -in ./keys/rsa -pubout -out ./keys/rsa.pub

.PHONY: swag
swag:
	swag init -g ./main.go --output docs/instalike

.PHONY: build
run:
	go build -o bin/ ./...
	cp ./.env bin/
	cp -r ./keys/ bin/

build:
	CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-s -w -extldflags=-static" -o bin/ ./...
	cp ./.env bin/
	cp -r ./keys/ bin/


.PHONY: postgres
postgres:
	docker run --rm -itd --network host \
		-e POSTGRES_USER=pusr \
		-e POSTGRES_PASSWORD=pusr_secret \
		-e POSTGRES_DB=instalike-data \
		-p 5432:5432 -v /postgres-instalike-data:/var/lib/postgresql/data \
		--name instalike-pg \
		postgres:latest

# Decided not to use migrations for now
# uncomment if/when I decide to go back to them
# .PHONY: migration
# migration:
# 	migrate create -ext sql -dir migrations -seq ${seq}

# .PHONY: up
# up:
# 	migrate -path migrations -database "postgresql://pusr:pusr_secret@localhost:5432/instalike-data?sslmode=disable" up

# .PHONY: down
# down:
# 	migrate -path migrations -database "postgresql://pusr:pusr_secret@localhost:5432/instalike-data?sslmode=disable" down 1