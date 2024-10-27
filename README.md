# ミッションのサンプル実装
様々なゲームシステムと密結合になりがちなミッション機能を疎結合にし、統一的に捉えられるように実装したサンプル


## ディレクトリ設計

- pkg/grpc
   - protoファイルをgo用にコンパイラした生成物ができる
- proto
   - protobufのファイルが入る。サブモジュール
- script
   - go用のprotobufコンパイラshell scriptが入る
- server
   - go言語のコード群が入る。

## アーキテクチャ
オニオンアーキテクチャを採用
https://amasuda.xyz/post/2023-01-12-pros-cons-ddd-and-golang/
https://qiita.com/little_hand_s/items/2040fba15d90b93fc124
https://zenn.dev/jy8752/articles/b7bea6802e9f02
![img.png](assets/img.png)

層の外から
Infrastructure, UI(Presentation) -> UseCase(Application Service) -> Domain

- アプリケーションは、独立したオブジェクト・モデルを中心に構築される
- 内側のレイヤーはインターフェースを定義。外側のレイヤーはインターフェースを実装
- 結合の方向は中心に向かっている
- すべてのアプリケーションのコアコードは、インフラストラクチャとは別にコンパイルして実行することができる


依存性の注入のため、registryを導入
https://moneyforward-dev.jp/entry/2021/03/08/go-test-mock/

Application Service層はUsecaseと命名する。
Protobufのserviceとかぶるし、単純にserviceという名前が低凝縮になりやすい

## Docker init

Composeコマンドを軸としてローカルのコンテナ環境を立ち上げます。
```
# 初回起動
docker compose build
docker compose up -d

# シャットダウン
docker compose down

# 2回目以降の起動
docker compose up -d
```