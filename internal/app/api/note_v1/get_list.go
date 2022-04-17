package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) GetList(ctx context.Context, in *desc.Empty) (*desc.GetListResponse, error) {
	fmt.Println("returned all notes")

	return &desc.GetListResponse{Results: []*desc.GetListResponse_Result{
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
