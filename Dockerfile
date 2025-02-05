FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
CMD [ "./main" ]