FROM golang:1.16

WORKDIR /app

ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

CMD ["./app"]

