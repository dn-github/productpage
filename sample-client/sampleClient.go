package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/dn-github/productpage/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	client := pb.NewProductPageServiceClient(conn)

	req := &pb.Book{
		Name: "The Book Thief",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := client.Product(ctx, req)
	if err != nil {
		log.Fatalf("error while calling gRPC: %v", err)
	}
	log.Printf("Response from Service: %v", res)
}
