package main

import (
	"context"
	"log"

	"grpc/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewAddServiceClient(conn)
	client.AddGarment(context.Background(), &proto.Item{Garment: "Shirt"})
}
