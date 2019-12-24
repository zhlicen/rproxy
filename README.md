# rproxy

rproxy supports transparent reverse proxy for multiple protocols(HTTP1.*/gRPC/gRPC-Web) with single entry port

## Protocols

| protocol                                     | desc         |
| -------------------------------------------- | ------------ |
| HTTP 1.*                                     |
| [gRPC](https://github.com/grpc)              | Native gRPC  |
| [gRPC-Web](https://github.com/grpc/grpc-web) | gRPC for web |
 

## Structure

![](./img/rproxy.png)


## Reference

rproxy use the following package:

| package                                                                  | desc                                      |
| ------------------------------------------------------------------------ | ----------------------------------------- |
| [soheilhy/cmux](https://github.com/soheilhy/cmux  )                                       | serve different services on the same port |
| [improbable-eng/grpc-web](https://github.com/improbable-eng/grpc-web/go/grpcweb) | in process grpc-web wrapper               |
| [mwitkow/grpc-proxy](https://github.com/mwitkow/grpc-proxy)                             | gRPC reverse proxy                        |

