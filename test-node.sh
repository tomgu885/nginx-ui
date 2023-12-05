#!/bin/env bash
## for test / 测试用
echo "build"

## stage 1 make files
dat=$(date +%Y%m%d%H%M)

GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME} -s -w" -o bin/nginx-node-${dat} cmd/actor/main.go

echo "built, start deploy node"
## stage 2, upload, restart cdn nodes// test 1 node
scp bin/nginx-node-${dat} cdn1:/data/nginx-node/
ssh cdn1 "ln -sfn /data/nginx-node/nginx-node-${dat} /data/nginx-node/nginx-node && systemctl restart nginx-node"

echo "finish"
