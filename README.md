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
Protobufのserviceとかぶるし、単純にserviceという名前が低凝縮になりやすい。

## migration
go-migrationを使用

https://zenn.dev/farstep/books/f74e6b76ea7456/viewer/4cd440

```shell
# migrationファイル作成
migrate create -ext sql -dir db/migrations -seq create_users_table

# migration実行
migrate --path db/migrations --database 'mysql://root:root@localhost:3306/centray-{env}-db' -verbose up
```

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


### ECSデプロイ
#### コンテナイメージをビルドしてECRにpushする
```
$ ./script/build-and-push-image.sh --tag dev1
Successfully pushed the image to ECR with tag 236f295
specify image 224735682946.dkr.ecr.ap-northeast-1.amazonaws.com/nornapi-app:236f295 in your task definition

Successfully pushed the image to ECR with tag latest
specify image 224735682946.dkr.ecr.ap-northeast-1.amazonaws.com/nornapi-app:latest in your task definition

Successfully pushed the image to ECR with tag dev4
specify image 224735682946.dkr.ecr.ap-northeast-1.amazonaws.com/nornapi-app:dev4 in your task definition
```
ECRリポジトリにpushされたイメージにはコミットハッシュ、latest、--tagで指定したものがタグとして付けられます。  
今後はこのタグを指定してデプロイすることを想定しています。

#### latestタグを利用してデプロイする
```
./script/deploy-latest.sh --env dev1
```
このコマンドではlatestタグのコンテナイメージを利用して以下の2点を実行しています。
* DBマイグレーションの実行
* ECSサービスのデプロイ

#### タグを指定してデプロイする
TODO: 未実装

## 認証

JWT
https://christina04.hatenablog.com/entry/2017/04/15/042646

config以下にRSA秘密鍵をおくことで著名と暗号化を行う。とりあえずトークンに含めるのはuserIdとuuid
本開発フェーズで、sessionIdをRedis管理し、そのsessionIdをJWTにしクライアントに渡すことでセキュリティレベルを高くしたい