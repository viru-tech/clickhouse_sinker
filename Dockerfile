FROM registry.lucky-team.pro/core/go.docker-images/debian:1.24 AS builder

ADD . /app
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN make build

ENTRYPOINT ["/app/bin/clickhouse_sinker"]
