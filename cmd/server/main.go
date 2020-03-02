package main

import (
	"log"
	"net"
	"transmitter/pkg/api"
	"transmitter/pkg/transmitter"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &transmitter.GRPCServer{}
	api.RegisterTransmitterServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
