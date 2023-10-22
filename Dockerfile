FROM golang:1.20-alpine AS build-env

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /tmp ./...

# Runtime image

FROM golang:1.20-alpine

WORKDIR /app
COPY --from=build-env /tmp/sqzsvc /app

CMD ["/app/sqzsvc"]