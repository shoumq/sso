FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o migrator ./cmd/migrator
RUN go build -o sso ./cmd/sso

CMD ["./migrator", "--storage-path=./storage/sso.db", "--migration-path=./migrations"]

CMD ["./sso", "--config=config/local.yaml"]