package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	pb "grpc-micro-demo-cli/proto"
	"time"
)

func main() {
	// New Service
	helloworld := micro.NewService(
		micro.Name("helloworld-client"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	// Initialise service
	helloworld.Init()

	ctx := context.TODO()
	ctx, _ = context.WithTimeout(ctx, 1*time.Minute)

	// Call the rpc
	service := pb.NewHelloworldService("helloworld", helloworld.Client())
	rsp, err := service.Call(ctx, &pb.Request{Name: "xie"})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Msg)

}
