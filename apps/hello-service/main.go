package main

import (
	"fmt"

	"net"
	"os"

	"github.com/dineshg13/goauto-play/libs/logging"
	"github.com/dineshg13/goauto-play/protos/hellopb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 8080

func main() {
	if err := realMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var logger *zap.Logger

func realMain() error {
	var portStr = fmt.Sprintf(":%d", port)
	var err error
	logger, err = logging.NewZapLogger()
	if err != nil {
		return err
	}
	lis, err := net.Listen("tcp", portStr)

	if err != nil {

		return err
	}

	logger.Info("Starting on port ", zap.Int("port", port))

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	hellopb.RegisterGreetingServiceServer(grpcServer, NewServer())
	if err = grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
