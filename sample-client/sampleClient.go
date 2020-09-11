package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/dn-github/productpage/pb"
)

func main() {
	fmt.Println("Enter QPS (0 for single run and exit): ")
	var qps int
	_, err := fmt.Scanf("%d", &qps)

	conn, err := grpc.Dial("192.168.99.109:32053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()
	client := pb.NewProductPageServiceClient(conn)

	req := &pb.Book{
		Name: "The Book Thief",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	sleepTime := time.Second/time.Duration(qps)

	for {
		res, err := client.Product(ctx, req)
		if err != nil {
			log.Fatalf("error while calling gRPC: %v", err)
		}
		log.Printf("Response from Service: %v", res)
		time.Sleep(sleepTime)
	}
}
