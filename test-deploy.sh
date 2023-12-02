#!/bin/env bash
## for test / 测试用
echo "build"
## stage 1 make files
dat=$(date +%Y%m%d%H%M)
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME} -s -w" -o bin/nginx-master-${dat} main.go
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.BuildTime=${BUILD_TIME} -s -w" -o bin/nginx-node-${dat} cmd/actor/main.go

#echo "built, start deploy node"
## stage 2, upload, restart cdn nodes// test 1 node
#scp bin/nginx-node-${dat} cdn1:/data/nginx-node/
#ssh cdn1 "ln -sfn /data/nginx-node/nginx-node-${dat} /data/nginx-node/nginx-node ; systemctl restart nginx-node"

echo "start deploy nginx-master"

scp bin/nginx-master-${dat} test:/data/nginx-master/
ssh test "ln -sfn /data/nginx-master/nginx-master-${dat} /data/nginx-master/nginx-master"
