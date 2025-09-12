FROM golang:1.24-alpine3.21

WORKDIR /go/src

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            graphviz

COPY ./src .

# go.modがあれば依存関係をインストール
RUN if [ -f ./go.mod ]; then \
      go install; \
    fi

# 開発用のライブラリをインストール
RUN go install github.com/air-verse/air@v1.62.0
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

EXPOSE 8080