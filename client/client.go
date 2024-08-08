package main

import (
	"context"
	"log"
	"net/http"

	pb "grpc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.HelloServiceClient

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client = pb.NewHelloServiceClient(conn)

	r := gin.Default()
	r.GET("/send-Message/:message", clientConnection)
	r.Run(":8080")
}

func clientConnection(c *gin.Context) {
	message := c.Param("message")

	req := &pb.HelloRequest{
		Name: message,
	}
	resp, err := client.SayHello(context.TODO(), req)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
	})

}
