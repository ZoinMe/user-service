FROM golang:1.22
WORKDIR /app
ADD . /app/
RUN go build -o ./out/zoinme-user-service .
EXPOSE 8080
ENTRYPOINT ["./out/zoinme-user-service"]