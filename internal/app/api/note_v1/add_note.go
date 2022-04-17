package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) AddNote(ctx context.Context, req *desc.AddNoteRequest) (*desc.AddNoteResponse, error) {
	fmt.Println("note added")

	return &desc.AddNoteResponse{
		Result: &desc.AddNoteResponse_Result{
			Id: 228,
		},
	}, nil
}
