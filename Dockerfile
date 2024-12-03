FROM golang:1.22.3 as build-deps

WORKDIR /usr/src/backend

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .
RUN go build /usr/src/backend/cmd/main.go

FROM alpine:3.19.1
WORKDIR /usr/src/app

ARG MODE
COPY --from=build-deps /usr/src/backend/run.sh run.sh
COPY --from=build-deps /usr/src/backend/main main
COPY --from=build-deps /usr/src/backend/configs/$MODE ./config

ARG MODULE
RUN chmod +x run.sh
RUN apk add --no-cache bash
RUN apk add --no-cache libc6-compat

ENV LOG_PATH=/logs/$MODULE.log

ENTRYPOINT ["./run.sh"]