# Build environment
# -----------------
FROM golang:1.17-alpine as build-env
WORKDIR /bookshelf

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /bookshelf/bin/app /bookshelf/

EXPOSE 8080
CMD ["/bookshelf/app"]
