# 管理方法

centray-api, centary-clientの中でサブモジュールとして管理

# protobuf

クライアントとサーバがAPI通信するためのインターフェースを定義
今回はgrpc通信のインターフェースを定義してく

grpcだと
- message
  - 基本的にEntityや、リクエスト、レスポンスの中身を定義
- service
  - grpc通信のルールをかく。メソッド名とか、streamの方向とかを定義


# ディレクトリ構造

```
.
├── README.md
├── Tools
│        ├── install.sh
│        ├── install_unity_env.sh
│        └── protocs.sh
├── src
│        ├── test.proto
│        └── test_service.proto
```

## Tools
ツール置き場

環境構築系や、コードの自動生成用Scriptなどの置き場

## proto ファイルの作成・運用方針について

- フラットに書く
  - ディレクトリ分けるの辛すぎるのがnornでの学び
- 命名は以下
  - messageは{entity_name}.proto
  - serviceは{service_name}_service.proto
- 依存方向は service -> message
  - つまり、serviceだけmessageをrequireできる


# 環境構築
## 共通
### Mac User
*前提*

- Homebrew はいれておいてください


1. `$ sh Tools/install.sh` を実行
2. 出力が `libprotoc 25.2` と出ていればOK

### For Windows
TBD

## Server

## Unity
`https://qiita.com/mattak/items/0c5e6065459f251b87c5` を参考に環境を構築します。

### UnityProject 側
基本的に導入済みになっているはずですが、新規PJでやる場合は以下をお願いします

1. https://qiita.com/akiojin/items/ac05392d97abb8797dcd を参考にScopedRegistortyに `UnityNuget` を追加
1. UPMから `Google.Protobuf (NuGet)` を追加
1. https://qiita.com/mattak/items/0c5e6065459f251b87c5#2-unity%E3%81%A7%E3%81%AE%E8%A8%AD%E5%AE%9A を見ながらYetAnotherHttpHandlerを導入する

### proto リポジトリ側

1. `$ cd Tools`
2. `$ sh install_unity_env.sh`


# コードの自動生成
## Server 側

go用。centray-apiで定義

```shell
#!/bin/bash

set -eu

protoc \
  --go_out=./pkg/grpc \
  --go_opt=paths=source_relative \
  --go-grpc_out=./pkg/grpc \
  --go-grpc_opt=paths=source_relative \
  -I./proto \
  ./proto/*.proto
```

## Unity 側

1. `$ sh Tools/protocs.sh`


# メモ
- テストの開始と終了はserver　streamで判断
- 100にんと接続中、特定の10にんにのみstreamする可能性あり
  - streamごとにidを発行し、それをインスタンスに格納し、stream_idsに保持 -> 送りたいstream_idで検索し、インスタンス経由でstream.sendする
- 二つイベント同時に走らせる
  - event_idとuser_idの二つくみでstreamに保存する？
- 識別子idは、APIサーバ側でuuidを発行して、それを最初に叩いてもらう
  - 認証の時も拡張性高い