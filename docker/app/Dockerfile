FROM golang:1.17-alpine
WORKDIR /bookshelf
RUN apk update && apk add --no-cache gcc musl-dev git bash
COPY . .
RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app
CMD ["/bookshelf/bin/app"]
EXPOSE 8080
