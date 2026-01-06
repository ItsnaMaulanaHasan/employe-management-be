FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o seeder ./cmd/seeder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/seeder .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/migrations ./migrations

RUN mkdir -p storage/reports

EXPOSE 8080

CMD ["./app"]
