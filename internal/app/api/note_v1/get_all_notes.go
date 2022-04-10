package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) GetAllNotes(ctx context.Context, in *desc.Empty) (*desc.GetAllNotesResponse, error) {
	fmt.Println("returned all notes")

	return &desc.GetAllNotesResponse{Results: []*desc.GetAllNotesResponse_Result{
		{
			Title:   "title1",
			Content: "content1",
		},
		{
			Title:   "title2",
			Content: "content2",
		},
	}}, nil
}
