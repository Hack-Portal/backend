# HackPortalBackend

## 動作環境

- docker
- docker-compose
- make
- go 1.20.6
- gin
- swag
- go-migrations
- swagger

## ローカルでの動かし方

1. 動作環境を整える
2. 任意の新規フォルダにて以下のコマンドを実行する

```bash
make postgresRun
make postgresStart
make resetdb
make migrateup
make serverRun
```

## memo

app.env を作って
DB_DRIVER=postgres
DB_SOURSE=DB接続用URI
を記述する必要あり
