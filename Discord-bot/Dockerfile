# ベースイメージ
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /go/src/app

# ホストのソースコードをコンテナにコピー
COPY . .

# MySQLのドライバをインストール
RUN go get -u github.com/go-sql-driver/mysql

# ポートを公開
EXPOSE 8080

# サーバーを実行
CMD ["go", "run", "main.go"]
