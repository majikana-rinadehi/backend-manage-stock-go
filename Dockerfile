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

RUN echo "running dev stage"

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.7
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install -v golang.org/x/tools/gopls@latest

# TODO: devcontainerで共通的に使うので、別の共通dockerfileとかにまとめたいよね。。。開発用のツールとかは
RUN apt-get update && apt-get install -y sudo
RUN sudo apt-get update && sudo apt-get install vim -y


COPY . .