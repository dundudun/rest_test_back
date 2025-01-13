FROM golang:1.23.4-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/app ./cmd/main

#RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
#RUN sqlc generate

EXPOSE 8080

CMD ["app"]