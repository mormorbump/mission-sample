#!/usr/bin/env bash

set -e

# --envパラメータの存在を確認
while [[ "$#" -gt 0 ]]; do
  case $1 in
    --env) AWS_ENV="$2"; shift ;;
    *) echo "Unknown parameter passed: $1"; exit 1 ;;
  esac
  shift
done

# AWS_ENVの存在を確認
if [ -z "$AWS_ENV" ]; then
  echo "AWS_ENV name not specified. Use --env to specify the environment.";
  exit 1
fi

# AWSリソースのNameタグの命名規則が環境ごとに揃っていることが前提
PROJECT_NAME="centray"
CLUSTER_NAME="${PROJECT_NAME}-${AWS_ENV}-cluster"
SERVICE_NAME="app-service"
MIGRATE_TASK_NAME="${PROJECT_NAME}-${AWS_ENV}-migrate"
REGION="ap-northeast-1"
SUBNET_TAG_VALUE="${PROJECT_NAME}-${AWS_ENV}-private-app-subnet-1a"
SG_TAG_VALUE="${PROJECT_NAME}-${AWS_ENV}-app-sg"

# ECSクラスターが存在するか確認
if ! aws ecs describe-clusters --clusters "${CLUSTER_NAME}" --region ${REGION} | grep "ACTIVE"; then
    echo "Specified environment does not exist or is not active in ECS."
    exit 1
fi

# Subnet IDを取得
SUBNET_IDS=$(aws ec2 describe-subnets \
  --filters "Name=tag:Name, Values=$SUBNET_TAG_VALUE" \
  --query "Subnets[*].SubnetId" \
  --output text)

# Security Group IDを取得
SG_ID=$(aws ec2 describe-security-groups \
  --filters "Name=tag:Name,Values=$SG_TAG_VALUE" \
  --query "SecurityGroups[*].GroupId" \
  --output text)


echo "Start Run Task"
aws ecs run-task \
  --cluster "${CLUSTER_NAME}" \
  --task-definition "${MIGRATE_TASK_NAME}" \
  --launch-type FARGATE \
  --network-configuration "awsvpcConfiguration={subnets=[$SUBNET_IDS],securityGroups=[$SG_ID]}"

echo "Start Update Service"
# ECSサービスの更新
aws ecs update-service \
  --cluster "${CLUSTER_NAME}" \
  --service ${SERVICE_NAME} \
  --force-new-deployment

echo "finish all"