### Generate proto user
```
Generate to micro user

api-gateway
$ protoc --proto_path=$GOPATH/src:./micro-user/pkg/pb \
    --micro_out=./micro-user/pkg/pb \
    --go_out=./micro-user/pkg/pb \
    --proto_path=./proto user.proto

Generate to api-gateway

api-gateway
$ protoc --proto_path=$GOPATH/src:./api-gateway/pkg/pb \
    --micro_out=./api-gateway/pkg/pb \
    --go_out=./api-gateway/pkg/pb \
    --proto_path=./proto user.proto
```

### Generate proto authorization
```
Generate to micro authorization

$ protoc --proto_path=$GOPATH/src:./micro-authorization/pb \
api-gateway
    --micro_out=./micro-authorization/pkg/pb \
    --go_out=./micro-authorization/pkg/pb \
    --proto_path=./proto authorization.proto

Generate to api-gateway

api-gateway
$ protoc --proto_path=$GOPATH/src:./api-gateway/pkg/pb \
    --micro_out=./api-gateway/pkg/pb \
    --go_out=./api-gateway/pkg/pb \
    --proto_path=./proto authorization.proto
```