package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.NoteInfo, error) {
	note, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	if note == nil && err == nil {
		return nil, status.Error(codes.NotFound, "there is no note with this id")
	}

	return note, nil
}
