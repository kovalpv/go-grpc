package transmitter

import (
	"context"
	"fmt"
	"transmitter/pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Send(ctx context.Context, req *api.SendRequest) (*api.SendResponse, error) {
	println(req.GetTime())
	timeline := req.GetTimeline()
	values := req.GetValues()
	for i, point_time := range timeline {
		value := values[i]
		fmt.Printf("%v => %v\n", point_time, value)
	}
	return &api.SendResponse{Success: true}, nil
}
