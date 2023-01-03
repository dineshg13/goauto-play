package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dineshg13/goauto-play/protos/hellopb"

	"go.uber.org/zap"
)

type Server struct {
	hellopb.UnimplementedGreetingServiceServer
}

func (s *Server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	if request.GetPerson() == nil {
		return &hellopb.HelloResponse{
			Code:         hellopb.Code_BAD_REQUEST,
			ErrorMessage: fmt.Sprintf("Person can't be empty"),
		}, nil
	}
	if request.GetPerson().Name == "" {
		return &hellopb.HelloResponse{
			Code:         hellopb.Code_BAD_NAME,
			ErrorMessage: fmt.Sprintf("Name can't be empty"),
		}, nil
	}
	select {
	case <-ctx.Done():
		logger.Error("context done", zap.Error(ctx.Err()))
		return nil, ctx.Err()
	case <-time.After(1 * time.Second):
		logger.Info("Time up sending response")

	}
	return &hellopb.HelloResponse{
		Code:            hellopb.Code_OK,
		ErrorMessage:    "",
		ResponseMessage: fmt.Sprintf("Hello %v", request.GetPerson().Name),
	}, nil
}

func NewServer() *Server {
	return &Server{}
}
