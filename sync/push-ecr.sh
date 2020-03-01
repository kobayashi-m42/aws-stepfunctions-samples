#!/usr/bin/env bash

if [[ "$1" = "" ]]; then
  echo  "AWSアカウントIDを第1引数に指定して下さい"
  exit 1
fi

ACCOUNT_ID="$1"
APP_NAME=dev-stepfunctions-sync

$(aws ecr get-login --no-include-email --region ap-northeast-1)

docker build -t ${ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/${APP_NAME}:latest -f Dockerfile .
docker push ${ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/${APP_NAME}:latest
