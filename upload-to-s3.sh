#!/bin/bash

set -e

# === 変数定義 ===
BUCKET_NAME="migration-tool-0331"
BINARY_NAME="migrate-tool"
SQL_FILE_PATH="./database/001_init.up.sql"
S3_SQL_KEY="database/001_init.up.sql"

echo "$(aws sts get-caller-identity)"

# === ① Goバイナリをビルド（Linux用）===
echo "🔨 ビルド中..."
GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} ./cmd/migrate

# === ② バイナリをS3へアップロード ===
echo "☁️ migrate-tool アップロード中..."
aws s3 cp ./${BINARY_NAME} s3://${BUCKET_NAME}/${BINARY_NAME}

# === ③ SQLファイルをS3へアップロード ===
echo "📄 SQLファイル アップロード中..."
aws s3 cp ${SQL_FILE_PATH} s3://${BUCKET_NAME}/${S3_SQL_KEY}

echo "✅ アップロード完了！"

# === ④ ローカルのバイナリを削除 ===
echo "🧹 ローカルのバイナリを削除中..."
rm -f ${BINARY_NAME}

echo "✅ 全処理完了！"
