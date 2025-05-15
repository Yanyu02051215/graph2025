# ビルド用ステージ
FROM golang:1.23 AS builder

WORKDIR /app

# 依存モジュールを先に取得（キャッシュ効率化）
COPY go.mod go.sum ./
RUN go mod download

# アプリ全体をコピーしてビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migrate-tool ./cmd/migrate2

# 実行用ステージ（超軽量）
FROM alpine:3.18

WORKDIR /app

# ビルド済みバイナリだけコピー
COPY --from=builder /app/server .
COPY --from=builder /app/migrate-tool .
COPY --from=builder /app/database ./database

# HTTP ポートを明示（Minikubeでのポートマッピング向け）
EXPOSE 8080

# コンテナ起動時の実行コマンド
CMD ["./server"]
