package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) RemoveNote(ctx context.Context, req *desc.RemoveNoteRequest) (*desc.Empty, error) {
	fmt.Printf("note with id=%v removed\n", req.GetId())

	return &desc.Empty{}, nil
}
