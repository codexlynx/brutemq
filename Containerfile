FROM golang:1.18.3-alpine3.16 AS builder

LABEL org.opencontainers.image.source=https://github.com/codexlynx/brutemq

WORKDIR /go/src/github.com/codexlynx/brutemq
RUN apk add gcc musl-dev --no-cache

COPY . .
RUN go build -buildmode=pie -ldflags '-linkmode external -extldflags "-static-pie"' -o /build/brutemq .

FROM scratch
COPY --from=builder /build/brutemq /
ENTRYPOINT ["/brutemq"]
