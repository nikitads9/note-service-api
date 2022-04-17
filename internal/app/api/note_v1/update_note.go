package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	fmt.Printf("edited note with id %v \n", req.Id)

	return &desc.Empty{}, nil
}
