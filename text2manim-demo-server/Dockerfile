# ビルドステージ
FROM golang:1.23.1-alpine AS builder

# 作業ディレクトリの設定
WORKDIR /app

# 依存関係ファイルのコピー
COPY go.mod go.sum ./

# 依存関係のダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# 実行ステージ
FROM alpine:latest

# 必要なCA証明書をインストール
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ビルドステージから実行可能ファイルをコピー
COPY --from=builder /app/main .

# アプリケーションの実行
CMD ["./main"]