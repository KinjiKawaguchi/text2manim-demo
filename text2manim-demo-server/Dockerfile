# ビルドステージ
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# 実行ステージ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

# .envファイルが存在する場合はコピー、存在しない場合は空ファイルを作成
COPY .env* ./.env 2>/dev/null || touch .env

CMD ["./main"]
