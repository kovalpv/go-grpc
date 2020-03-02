package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"transmitter/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	address := os.Args[1]

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal((err))
	}

	time := "2019-01-01T11:00:00+00"
	timeline := []string{"2019-01-01T10:00:00+00", "2019-01-01T10:30:00+00"}
	values := []float32{2.0, 5.0}

	c := api.NewTransmitterClient(conn)
	res, err := c.Send(context.Background(), &api.SendRequest{Time: time, Timeline: timeline, Values: values})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("message sended is %v", res.GetSuccess())
	println()
}
