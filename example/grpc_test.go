package main

import (
	"context"
	"fmt"
	"github.com/zhlicen/rproxy/example/proto"
	"google.golang.org/grpc"
	"testing"
)

func Test_grpcReq(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := proto.NewTestServiceClient(conn)
	req := proto.Req{Req: "hello backend!!!!"}
	rsp, err := cli.TestRequest(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func BenchmarkGRPC(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	for n := 0; n < b.N; n++ {
		cli := proto.NewTestServiceClient(conn)
		req := proto.Req{Req: "hello backend!!!!"}
		_, err := cli.TestRequest(context.Background(), &req)
		if err != nil {
			panic(err)
		}
		//fmt.Println(rsp)
	}
}
