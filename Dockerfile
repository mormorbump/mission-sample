# AWS CodeBuild の Rate Limit に引っかからないようにするために ECR Public Gallery から取得する
FROM public.ecr.aws/docker/library/golang:1.23.2-alpine AS builder

ENV ROOT=/app
WORKDIR ${ROOT}

RUN apk upgrade && apk add git
COPY server server
COPY pkg pkg
COPY go.mod go.sum ./
RUN go mod download

# 冗長性max && 重大度をinfoにしてできるだけログを出すようにする
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
ENV GRPC_GO_LOG_SEVERITY_LEVEL=info

RUN go build server/main.go

FROM alpine AS runner

ENV ROOT=/app
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
ENV GRPC_GO_LOG_SEVERITY_LEVEL=info
WORKDIR ${ROOT}

# --from=0: ファイルをコピーする元のステージを指定
# こうすることで最終的なイメージを最小化
# RUN go build server/main.goは/app以下に生成するのでこれで良い。
COPY --from=builder /app/main main
CMD ["/app/main"]

FROM alpine AS deploy

ENV ROOT=/app
WORKDIR ${ROOT}

COPY --from=builder /app/main main
CMD ["/app/main"]


# migration用
FROM golang:1.22.1-alpine AS migrate

RUN apk add --no-cache git mysql-client && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# マイグレーションファイルのコピー
COPY server/db/migrations /app/migrations

# マイグレーションを実行するスクリプトの追加
COPY script/image-migrate.sh /app/script/image-migrate.sh
RUN chmod +x /app/script/image-migrate.sh

# migrate.shを実行
CMD ["/app/script/image-migrate.sh"]


# 開発ステージ (live-reload用)
FROM golang:1.23.2-alpine AS dev-live-reload

RUN go install github.com/cosmtrek/air@v1.49.0

# compose.yamlで`.:/app`のマウント設定が存在することを前提とする
WORKDIR /app
CMD ["air", "-c", ".air.toml"]
