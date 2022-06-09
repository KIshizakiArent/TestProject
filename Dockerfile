FROM golang:1.18.3

# コンテナの作業ディレクトリにローカルのファイルをコピー
WORKDIR /app
COPY . /app

# 必要なパッケージをインストール
RUN go mod tidy

# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
CMD ["air"]