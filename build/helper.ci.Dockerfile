FROM golang:1.16

ARG CI_JOB_TOKEN

ENV GOPRIVATE github.com/Harzu/*
ENV GO111MODULE on

RUN set -e \
    && mkdir -p .go/bin \
    && git config --global url."https://Harzu:${CI_JOB_TOKEN}@github.com".insteadOf "https://github.com" \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar -xvz \
    && mv ./migrate.linux-amd64 /bin/migrate \
    && wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.38.0 \
    && mv ./bin/golangci-lint /bin/golangci-lint

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT []
