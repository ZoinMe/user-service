FROM golang:1.21.6-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .
EXPOSE 8080
CMD ["./main"]
