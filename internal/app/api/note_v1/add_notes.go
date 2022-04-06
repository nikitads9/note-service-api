package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) AddNotes(ctx context.Context, req *desc.AddNotesRequest) (*desc.AddNotesResponse, error) {
	results := []Results{}
	for i:= 0; i <len(req.Notes); i++ {
		ids = append(ids, int64(i))
	}
	fmt.Println("added multiple entries")
	return &desc.AddNotesResponse{
		Results: &desc.AddNotesResponse_Result{
			Id: ids,
		},
	}, nil
}