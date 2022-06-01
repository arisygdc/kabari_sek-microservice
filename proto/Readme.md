### Generate proto user
```
Generate to micro user

$ protoc --proto_path=$GOPATH/src:./micro-user/pb --micro_out=./micro-user/pb --go_out=./micro-user/pb --proto_path=./proto user.proto

Generate to api-gateway

$ protoc --proto_path=$GOPATH/src:./api-gateway/pb --micro_out=./api-gateway/pb --go_out=./api-gateway/pb --proto_path=./proto user.proto
```