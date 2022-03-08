package main

import (
	"fmt"
	"log"

	"github.com/Fkhalilullin/go-grpc-api/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	log.Printf("Created client %f", c)
}
