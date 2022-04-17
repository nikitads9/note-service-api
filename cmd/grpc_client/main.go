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
	//nolint
	con, err := grpc.Dial(grpcAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer con.Close()

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

	addedID, err := client.MultiAdd(ctx, &pb.MultiAddRequest{
		Notes: []*pb.MultiAddRequest_Notes{
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
	if err != nil {
		log.Printf("failed to remove note: %v\n", err)
	}

	fmt.Printf("IDs: %v", addedID.GetResult().Count)

	_, err = client.GetNote(ctx, &pb.GetNoteRequest{
		Id: 0,
	})
	if err != nil {
		log.Printf("failed to get note: %v\n", err)
	}

	_, err = client.GetList(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("failed to get all notes: %v\n", err)
	}

	_, err = client.UpdateNote(ctx, &pb.UpdateNoteRequest{
		Id:      0,
		Title:   "newtitle",
		Content: "newcontent",
	})
}
