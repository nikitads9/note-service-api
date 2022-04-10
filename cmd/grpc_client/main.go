package main

import (
	"context"
	"fmt"

	"log"

	"github.com/nikitads9/note-service-api/pkg/note_api"
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

	addedID, err2 := client.MultiAdd(ctx, &pb.MultiAddRequest{
		Notes: []*note_api.MultiAddRequest_Notes{
			{
				Title:   "title1",
				Content: "ffdsjfdjf",
			},
			{
				Title:   "title2",
				Content: "sometext",
			},
			{
				Title:   "title3",
				Content: "more text",
			},
		},
	})
	if err2 != nil {
		log.Printf("failed to remove note: %v\n", err)
	}
	fmt.Printf("IDs: %v", addedID.GetResults().Id)

	_, err = client.GetNote(ctx, &pb.GetNoteRequest{
		Id: 0,
	})
	if err != nil {
		log.Printf("failed to get note: %v\n", err)
	}

	_, err = client.GetAllNotes(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("failed to get all notes: %v\n", err)
	}

	_, err = client.EditNote(ctx, &pb.EditNoteRequest{
		Id: 0,
	})
}
