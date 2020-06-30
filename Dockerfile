FROM golang:1.14 AS builder

COPY . /src
WORKDIR /src

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest AS production

WORKDIR /root/

COPY --from=builder /src/app .
COPY --from=builder /src/migrations ./migrations

CMD ["./app server"]