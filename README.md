# HackPortalBackend

## 動作環境

### 必須環境

- docker
- go 1.20
- go-migrations

### 推奨環境

- docker
- go 1.20.7
- go-migrations
- gin
- swag
- make

## ローカルでのセットアップ

1. 動作環境を整える
2. 任意の新規フォルダにて以下のコマンドを実行して、リポジトリをクローンする

```bash
git clone https://github.com/Hack-Hack-geek-Vol6/backend
```

3. 以下のコマンドを実行する  

```bash
make postgresRun
make postgresStart
make resetdb
make migrateup
```

4. app.env を記述する

## memo

app.env を作って
DB_DRIVER=postgres
DB_SOURSE=DB接続用URI
を記述する必要あり
