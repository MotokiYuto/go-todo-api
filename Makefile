# Goプロジェクトの初期化
init:
	go mod init go-hello-world

# 依存関係の整理
tidy:
	go mod tidy

# Goの起動
run:
	go run ./golang/src/main.go

# Docker Composeを使用してコンテナを立ち上げる
up:
	docker compose up -d

# Docker Composeを使用してコンテナを停止する
stop:
	docker compose stop

# Docker Composeを使用してコンテナを削除する
down:
	docker compose down

# Postgresコンテナに入る
postgres:
	docker exec -it postgres /bin/sh

go:
	docker exec -it go /bin/sh
# コンテナに入った後のコマンド
# psql -h localhost -U postgres -d postgres
