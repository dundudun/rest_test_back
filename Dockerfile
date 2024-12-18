FROM go:1.23.3

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build ./cmd/main/main.go -v -o /usr/local/bin/app ./...

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate

EXPOSE 8080

CMD ["app"]