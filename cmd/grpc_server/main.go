package main

import (
	"log"
	"net"

	"github.com/nikitads9/note-service-api/internal/app/api/note_v1"
	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
)

const grpcAdress = ":50051"

func main() {
	list, err := net.Listen("tcp", grpcAdress)
	if err != nil {
		log.Fatalf("failed to map port: %v", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterNoteV1Server(s, &note_v1.Implementation{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to process gRPC server: %v", err.Error())
	}

}
