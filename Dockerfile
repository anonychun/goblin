FROM golang:1.25.0 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o bin/server cmd/server/main.go
RUN go build -o bin/db cmd/db/main.go

FROM debian:trixie-slim
WORKDIR /app

RUN apt-get update -qq && \
	apt-get install --no-install-recommends -y curl wget telnet htop vim tmux tini postgresql-client

COPY --from=build /app/bin ./bin

ENTRYPOINT ["/usr/bin/tini", "--", "/app/bin/docker-entrypoint"]
CMD ["./bin/server", "start"]
