package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
)

const grpcAdress = "localhost:50051"

func main() {
	ctx := context.Background()
	//nolint
	con, err := grpc.Dial(grpcAdress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err.Error())
	}
	defer con.Close()

	client := pb.NewNoteV1Client(con)

	var res *pb.AddNoteResponse
	res, err = client.AddNote(ctx, &pb.AddNoteRequest{
		Note: &pb.Notes{
			Title:   "title1",
			Content: "fhdshjdsgd",
		},
	})
	if err != nil {
		log.Printf("failed to add note: %v\n", err.Error())
	}

	fmt.Println("note with id", res.GetId(), "added")

	var addedID *pb.MultiAddResponse

	addedID, err = client.MultiAdd(ctx, &pb.MultiAddRequest{
		Notes: []*pb.Notes{
			{
				Title:   "title11",
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
		log.Printf("failed to add notes: %v\n", err.Error())
	}
	fmt.Printf("added %v IDs\n", addedID.GetCount())

	_, err = client.RemoveNote(ctx, &pb.RemoveNoteRequest{Id: int64(4)})
	if err != nil {
		log.Printf("failed to remove note: %v\n", err.Error())
	}

	notes, err := client.GetList(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("failed to get all notes: %v\n", err.Error())
	}
	fmt.Printf("%v\n", notes.GetNoteInfo())

	_, err = client.UpdateNote(ctx, &pb.UpdateNoteRequest{
		Id: 2,
		Title: &wrapper.StringValue{
			Value: "newtitle",
		},
		Content: &wrapper.StringValue{
			Value: "newcontent UPDATED",
		},
	})
	if err != nil {
		log.Printf("failed to update a note: %v\n", err.Error())
	}

	note, err := client.GetNote(ctx, &pb.GetNoteRequest{
		Id: 2,
	})
	if err != nil {
		log.Printf("failed to get note: %v\n", err.Error())
	}
	fmt.Printf("%v\n", note)
}
