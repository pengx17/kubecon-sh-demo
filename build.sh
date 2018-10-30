#!/bin/bash

rm -rf dist
echo ">>> Building k8sswag:build"
docker build -t pengxiao/k8sswag:build -f ./dockerfiles/Dockerfile.build .
docker container create --name k8sswag-build-extract pengxiao/k8sswag:build
docker container cp k8sswag-build-extract:/dist ./dist
docker container rm k8sswag-build-extract

echo ">>> Building k8sswag:$1"
docker build --no-cache -t pengxiao/k8sswag:$1 -f ./dockerfiles/Dockerfile .
rm -rf dist
