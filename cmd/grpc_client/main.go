package main

import (
	"context"
	"fmt"

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
	res, err := client.AddNote(ctx, &pb.AddNoteRequest{
		Title:   "title1",
		Content: "fhdshjdsgd",
	})
	if err != nil {
		log.Printf("failed to add note: %v\n", err)
	}

	fmt.Println("note id =", res.GetResult().GetId())

	_, err = client.RemoveNote(ctx, &pb.RemoveNoteRequest{Id: int64(228)})
	if err != nil {
		log.Printf("failed to remove note: %v\n", err)
	}

}
