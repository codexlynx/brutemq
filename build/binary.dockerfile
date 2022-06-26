FROM golang:1.18.3-buster AS builder
WORKDIR /go/src/github.com/codexlynx/brutemq

RUN apt-get update -y \
    && apt-get install gcc-arm-linux-gnueabi -y \
    && mkdir -p /build/dist/

COPY . .

# amd64 build
RUN go build -o /build/dist/brutemq_amd64 ./

# arm build
RUN GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o /build/dist/brutemq_arm ./

FROM scratch AS binary
COPY --from=builder /build/dist /
