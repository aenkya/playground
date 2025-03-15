FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o playground main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/playground .

EXPOSE 8080

CMD ["./playground"]
