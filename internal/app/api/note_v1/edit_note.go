package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) EditNote(ctx context.Context, req *desc.EditNoteRequest) (*desc.Empty, error) {
	fmt.Printf("edited note with id %v \n", req.Id)

	return &desc.Empty{}, nil
}
