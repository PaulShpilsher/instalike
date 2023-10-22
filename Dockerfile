#################
# build image
#################
FROM golang:1.21.3-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify


# Build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags=-static" -o /tmp ./...


#################
# Runtime image
#################
FROM scratch

WORKDIR /app

# Copy keys and .env
COPY .env.docker ./.env
COPY keys/ ./
# Copy our static executable.
COPY --from=builder /tmp/instalike /app

# Run
ENTRYPOINT ["/app/instalike"]
