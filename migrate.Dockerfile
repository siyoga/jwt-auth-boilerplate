FROM golang:1.22.3

WORKDIR /usr/src/app

RUN go install github.com/pressly/goose/cmd/goose@latest
COPY migrations .

ENTRYPOINT [ "bash", "-c", "goose postgres 'host=db port=5432 user=admin password=1234 dbname=main sslmode=disable' up" ]