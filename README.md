# HackPortalBackend

## 動作環境

### 必須環境

- docker
- docker-compose
- go 1.22.2

## ローカルでのセットアップ

1. 動作環境を整える
2. 任意の新規フォルダにて以下のコマンドを実行して、リポジトリをクローンする

```bash
git clone https://github.com/Hack-Hack-geek-Vol6/backend
```

3. 以下のコマンドを実行する  

```bash
make rund #docker-compose経由で起動する
make seed #seedを投入する
```

4. app.env を記述する

## memo

app.env を作り
DB_DRIVER=postgres
DB_SOURSE=DB接続用URI
を記述する必要あり
