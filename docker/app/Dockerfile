# Build environment
# -----------------
FROM golang:1.17-alpine as build-env
WORKDIR /bookshelf

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /bookshelf/bin/app /bookshelf/
COPY --from=build-env /bookshelf/bin/migrate /bookshelf/
COPY --from=build-env /bookshelf/migrations /bookshelf/migrations

COPY --from=build-env /bookshelf/docker/app/bin /usr/local/bin/bookshelf/
RUN chmod +x /usr/local/bin/bookshelf/*

EXPOSE 8080
CMD ["/bookshelf/app"]
