FROM golang:1.18.1-alpine3.15

# アップデートとgit追加
RUN apk update && apk add git

# alpineなので開発に必要なパッケージをインストール
# 無いとテストが動かない
# https://qiita.com/KEINOS/items/fd6a299961e3b8f3864f
RUN apk add --no-cache alpine-sdk build-base


# appディレクトリ作成
RUN mkdir /go/src/app

ADD src/ /go/src/app

COPY go.mod /go/src/app/go.mod
COPY go.sum /go/src/app/go.sum

# ワーキングディレクトリの指定
WORKDIR /go/src/app

RUN go mod tidy
RUN go mod download

# Goの専用ツールをダウンロード
# https://zenn.dev/bun913/articles/f0a6c6177a4716
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls

# サーバー起動
# RUNだとイメージ開始時、CMDだとコンテナ開始時、ENTRYPOINTはコンテナ作成時
# CMD ["go", "run", "main.go"]