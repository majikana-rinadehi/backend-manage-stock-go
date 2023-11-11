# ビルド環境
FROM golang:1.19.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /backend-manage-stock-go

# 本番環境
FROM gcr.io/distroless/static  AS deploy

COPY --from=builder /backend-manage-stock-go /backend-manage-stock-go

EXPOSE 8080

ENTRYPOINT [ "/backend-manage-stock-go" ]

# 開発環境
FROM golang:1.19.2 AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.7

COPY . .