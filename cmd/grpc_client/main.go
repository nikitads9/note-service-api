package main

import (
	"context"

	"log"

	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
)

const grpcAdress = "localhost:50051"

func main() {
	ctx := context.Background()
	con, err := grpc.Dial(grpcAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error)
	}
	defer func() {
		err = con.Close()
		if err != nil {
			log.Fatalf("failed to close connection")
		}
	}()

	client := pb.NewNoteV1Client(con)
	_, err = client.AddNote(ctx, &pb.AddNoteRequest{
		Title:   "title1",
		Content: "fhdshjdsgd",
	})
	if err != nil {
		log.Printf("unimplemented API handler: %v", err)

	}

}
