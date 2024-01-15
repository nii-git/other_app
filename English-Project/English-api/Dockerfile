# ベースイメージの指定
FROM --platform=linux/x86_64 golang:latest
ARG ENV="local" 
ENV ENV ${ENV}

# 作業ディレクトリの設定
WORKDIR /api/
COPY go.mod go.sum /api/

# ソースコードをコピー
COPY . /api/

# ビルド
RUN go build -o main .

# アプリケーションの起動
CMD ["./main"]