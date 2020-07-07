package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/bajubullet/grpcdemo/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	msg := &chat.Message{Body: "Hello from client"}

	resp, err := c.SayHello(context.Background(), msg)
	if err != nil {
		log.Fatalf("Can't say hello: %s", err)
	}
	log.Printf("Response from server: %s", resp.Body)
}
