package main

import (
	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcAdress = ":50051"

func main() {
	list, err := net.Listen("tcp", grpcAdress)
	if err != nil {
		log.Fatalf("failed to map port: %v", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterNoteV1Server(s, &pb.UnimplementedNoteV1Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to process gRPC server: %v", err.Error())
	}
	

}
