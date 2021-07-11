FROM golang:1.16.5-buster AS builder

WORKDIR /opt/web-crawler

COPY ./go.mod   ./go.mod
COPY ./go.sum   ./go.sum
COPY ./vendor   ./vendor
COPY ./prebuild ./prebuild

RUN go build -v ./vendor/...
RUN go build -v -x -o /tmp/prebuild ./prebuild

COPY ./ ./

RUN go build -v ./cmd/web-crawler

FROM debian:buster

COPY --from=builder /opt/web-crawler/web-crawler /opt/web-crawler

CMD ["/opt/web-crawler"]
