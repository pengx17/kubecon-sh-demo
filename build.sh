#!/bin/bash

rm -rf dist
echo ">>> Building kubecon-sh-demo:build"
docker build -t pengxiao/kubecon-demo:build -f ./dockerfiles/Dockerfile.build .
docker container create --name kubecon-demo-build-extract pengxiao/kubecon-demo:build
docker container cp kubecon-demo-build-extract:/dist ./dist
docker container rm kubecon-demo-build-extract

echo ">>> Building kubecon-sh-demo:$1"
docker build --no-cache -t pengxiao/kubecon-demo:$1 -f ./dockerfiles/Dockerfile .
rm -rf dist
