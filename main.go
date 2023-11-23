package main

import (
	"calculator-unary-grpc-server/calculator"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 50051

type server struct {
	calculator.UnimplementedCalculateServer
}

func (s *server) Execute(ctx context.Context, in *calculator.CalculateRequest) (*calculator.CalculateResponse, error) {
	switch in.Operation {
	case calculator.Operation_ADD:
		return &calculator.CalculateResponse{
			Result: in.GetLhs() + in.GetRhs(),
		}, nil
	case calculator.Operation_SUB:
		return &calculator.CalculateResponse{
			Result: in.GetLhs() - in.GetRhs(),
		}, nil
	case calculator.Operation_MUL:
		return &calculator.CalculateResponse{
			Result: in.GetLhs() * in.GetRhs(),
		}, nil
	case calculator.Operation_DIV:
		return &calculator.CalculateResponse{
			Result: in.GetLhs() / in.GetRhs(),
		}, nil
	default:
		return &calculator.CalculateResponse{
			Result: 0,
		}, errors.New("invalid operation")
	}
}
func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("faild to listen: %v", err.Error())
	}

	s := grpc.NewServer()
	calculator.RegisterCalculateServer(s, &server{})

	log.Printf("server listen on %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("faild to serve : %v", err)
	}
}
