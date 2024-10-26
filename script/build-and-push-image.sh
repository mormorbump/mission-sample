#!/usr/bin/env bash

set -eu

# --tagパラメータの存在を確認
# 処理されていない引数があるかチェック
while [[ "$#" -gt 0 ]]; do
  case $1 in
    --tag) OPTION_TAG="$2" shift;;
    *) echo "Unknown parameter passed; $1"; exit 1 ;;
  esac
  shift
done

ACCOUNT_ID="637423435007"
REGION="ap-northeast-1"
ECR_ROOT="${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com"
APP_NAME="centrayapi-app"
MIGRATE_NAME="centrayapi-migrate"
ECR_APP_REPO="${ECR_ROOT}/${APP_NAME}"
ECR_MIGRATE_REPO="${ECR_ROOT}/${MIGRATE_NAME}"
LATEST_TAG="latest"
COMMIT_HASH=$(git rev-parse --short HEAD) # 最新のコミットアッシュを短くして取得

BUILD_TARGET="deploy"
SETUP_BUILD_TARGET="setup"

# 認証トークンを取得し、レジストリに対して Docker クライアントを認証する
aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR_ROOT}

##################appイメージ##################
# latest, commit hash のタグを付けてでイメージをビルドする
docker build \
  -t "${ECR_APP_REPO}:${COMMIT_HASH}" \
  --platform=linux/amd64 \
  --target="${BUILD_TARGET}" .

echo "push ${BUILD_TARGET}"
# ECR にイメージを push する
docker push "${ECR_APP_REPO}:${COMMIT_HASH}" > /dev/null
echo "Successfully pushed the image to ECR with tag ${COMMIT_HASH}"
echo "specify image ${ECR_APP_REPO}:${COMMIT_HASH} in your task definition"

# latestタグを付けてpushする
docker tag "${ECR_APP_REPO}:${COMMIT_HASH}" "${ECR_APP_REPO}:${LATEST_TAG}"
docker push "${ECR_APP_REPO}:${LATEST_TAG}" > /dev/null

echo "Successfully pushed the image to ECR with tag ${LATEST_TAG}"
echo "specify image ${ECR_APP_REPO}:${LATEST_TAG} in your task definition"

##################setupイメージ##################
docker build \
  -t "${ECR_MIGRATE_REPO}:${COMMIT_HASH}" \
  --platform=linux/amd64 \
  --target="${BUILD_TARGET}" .

echo "push ${BUILD_TARGET}"
# ECR にイメージを push する
docker push "${ECR_MIGRATE_REPO}:${COMMIT_HASH}" > /dev/null
echo "Successfully pushed the image to ECR with tag ${COMMIT_HASH}"
echo "specify image ${ECR_MIGRATE_REPO}:${COMMIT_HASH} in your task definition"

# latestタグを付けてpushする
docker tag "${ECR_MIGRATE_REPO}:${COMMIT_HASH}" "${ECR_MIGRATE_REPO}:${LATEST_TAG}"
docker push "${ECR_MIGRATE_REPO}:${LATEST_TAG}" > /dev/null

echo "Successfully pushed the image to ECR with tag ${LATEST_TAG}"
echo "specify image ${ECR_MIGRATE_REPO}:${LATEST_TAG} in your task definition"


# 追加のタグを指定してpushする(dev1など)
if [ ! -z "${OPTION_TAG+x}" ]; then
  # app task
  docker tag "${ECR_APP_REPO}:${COMMIT_HASH}" "${ECR_APP_REPO}:${OPTION_TAG}"
  docker push "${ECR_APP_REPO}:${OPTION_TAG}" > /dev/null
  echo "Successfully pushed the image to ECR with tag ${OPTION_TAG}"
  echo "specify image ${ECR_APP_REPO}:${OPTION_TAG} in your task definition"
  # migrate task
  docker tag "${ECR_MIGRATE_REPO}:${COMMIT_HASH}" "${ECR_MIGRATE_REPO}:${OPTION_TAG}"
    docker push "${ECR_MIGRATE_REPO}:${OPTION_TAG}" > /dev/null
    echo "Successfully pushed the image to ECR with tag ${OPTION_TAG}"
    echo "specify image ${ECR_MIGRATE_REPO}:${OPTION_TAG} in your task definition"
fi
