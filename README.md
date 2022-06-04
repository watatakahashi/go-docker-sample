# go-docker-sample

Go+PostgreSQLのDocker開発環境

## 環境構築

dockerを立ち上げる

```sh
docker-compose up -d
```

その後、コンテナに入り以下を実行

```sh
go run src/main.go
```

## テスト

カレントディレクトリ以下のテストをすべて実行する
```
go test -v ./...
```

## sqlcを使ったテーブルアクセス

sqlcで自動生成されたコードは`/sqlc`で管理する。原則としてRepositoryとテストからのみ使う。

sqcl導入後、以下のコマンドで生成できる。
```sh
sqlc generate
```

詳細は以下を参考にする
https://github.com/kyleconroy/sqlc

## API定義書のコード生成

openapi-generatorをインストールし、下記を実行\
// TODO:環境に依存しないインストール

```sh
cd openapi
mkdir output
openapi-generator generate -i v1.yaml -g go-gin-server -o ./output/
```

参考
https://openapi-generator.tech/
