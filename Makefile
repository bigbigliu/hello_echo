GOPATH:=$(shell go env GOPATH)
.PHONY : build

docker-build:
	docker build -t registry.cn-zhangjiakou.aliyuncs.com/develop_bigbigliu/hello_echo:v2 .

docker-push:
	docker push registry.cn-zhangjiakou.aliyuncs.com/develop_bigbigliu/hello_echo:v2