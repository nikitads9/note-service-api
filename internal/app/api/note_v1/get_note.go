package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {

	fmt.Printf("requested note with id %v \n", req.Id)
	return &desc.GetNoteResponse{
		Title:   "sometitle",
		Content: "some content",
	}, nil
}
