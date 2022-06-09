FROM golang:1.18-bullseye

# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app

# 必要なパッケージをインストール
RUN go mod tidy

# 起動
CMD ["go", "run", "/app/main.go"]