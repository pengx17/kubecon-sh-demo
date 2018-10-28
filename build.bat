REM Makes sure previous build is removed
RMDIR /s /q dist

ECHO Building kubecon-sh-demo:build
docker build -t pengxiao/kubecon-demo:build -f .\dockerfiles\Dockerfile.build .
docker container create --name kubecon-demo-build-extract pengxiao/kubecon-demo:build
docker container cp kubecon-demo-build-extract:/dist .\dist
docker container rm kubecon-demo-build-extract

ECHO Building kubecon-sh-demo:latest
docker build --no-cache -t pengxiao/kubecon-demo:latest -f .\dockerfiles\Dockerfile .

RMDIR /s /q dist
