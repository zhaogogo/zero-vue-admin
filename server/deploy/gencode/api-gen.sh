#!/usr/bin/env bash
protoc-go-inject-tag --input sysuser.pb.go

进入到usercenter/api/doc目录
goctl api go -api main.api --dir ../ --style goZero

进入到usercenter/rpc/system/doc
goctl rpc protoc sysuser.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero

系统用户model生成rpc目录操作
goctl model mysql datasource --url="root:zhaO..123@tcp(127.0.0.1:3306)/usercenter" --table="sys*" -d ./model/sysuser


