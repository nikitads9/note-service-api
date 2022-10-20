package note_v1

import (
	"context"

	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	note, err := i.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetNoteResponse{
		Id: note.Id,
		Note: &desc.Notes{
			Title:   note.Title.String,
			Content: note.Content.String,
		},
	}, nil
}
