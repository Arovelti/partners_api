FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o cmd/main .

FROM alpine:3.14
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/cmd/main .

ENV DATABASE_URL=postgres://user:password@db:5432/partner_db?sslmode=disable
ENV PORT=8080

CMD ["./cmd/main"]
