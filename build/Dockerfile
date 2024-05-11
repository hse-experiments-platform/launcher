FROM golang:1.22 AS builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/launcher ./cmd/launcher/main.go

FROM python:3.11-alpine

RUN apk update
RUN apk add make automake gcc g++ subversion python3-dev
RUN pip install --no-cache-dir pybind11 numpy requests pandas fire
RUN pip install -U  scikit-learn

WORKDIR /build

COPY configs /build/configs
COPY swagger /build/swagger
COPY scripts /build/scripts
COPY --from=builder /build/bin/launcher /build/bin/launcher

CMD ["./bin/launcher"]
