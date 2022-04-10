package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) MultiAdd(ctx context.Context, req *desc.MultiAddRequest) (*desc.MultiAddResponse, error) {
	ids := []int64{}
	for i := 0; i < len(req.Notes); i++ {
		ids = append(ids, int64(i))
	}
	fmt.Println("added multiple entries")
	return &desc.MultiAddResponse{
		Results: &desc.MultiAddResponse_Result{
			Id: ids,
		},
	}, nil
}
