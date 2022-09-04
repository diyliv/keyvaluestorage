package server

import (
	"log"
	"net"
	"os"
	"os/signal"

	grpcservices "github.com/diyliv/keyvaluestorage/internal/services/grpc"
	"github.com/diyliv/keyvaluestorage/proto/keyval"
	"google.golang.org/grpc"
)

func (s *server) RunGrpc() {
	log.Printf("Starting grpc server on port :50051")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Printf("Error while starting gRPC server: %v\n", err)
	}

	grpcServices := grpcservices.NewGrpcService(s.storage)

	serv := grpc.NewServer()
	keyval.RegisterKeyValueServer(serv, grpcServices)

	go func() {
		if err := serv.Serve(lis); err != nil {
			log.Printf("Error while serving: %v\n", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}
