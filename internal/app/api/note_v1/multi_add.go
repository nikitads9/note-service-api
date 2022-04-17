package note_v1

import (
	"context"
	"fmt"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) MultiAdd(ctx context.Context, req *desc.MultiAddRequest) (*desc.MultiAddResponse, error) {
	quantity := int64(len(req.Notes))

	fmt.Println("added multiple entries")

	return &desc.MultiAddResponse{
		Result: &desc.MultiAddResponse_Result{
			Count: quantity,
		},
	}, nil
}
