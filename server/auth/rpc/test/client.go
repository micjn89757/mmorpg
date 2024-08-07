package main

import (
	"context"
	"fmt"
	"server/common/idl/gen/auth"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	
	Client := auth.NewAuthClient(conn)

	// 执行RPC调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	r, err := Client.CheckToken(ctx, &auth.CheckTokenReq{
		Token: "01912b03-89f1-7c1e-89ee-f365d8936ed9",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("get reply:", r.GetData().Account)

}