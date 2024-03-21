FROM golang:1.22 AS builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/launcher ./cmd/launcher/main.go

FROM alpine:latest

WORKDIR /build

COPY configs /build/configs
COPY swagger /build/swagger
COPY --from=builder /build/bin/launcher /build/bin/launcher

CMD ["./bin/launcher"]
