package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/app/model"
)

func (s *Service) MultiAdd(ctx context.Context, notes []*model.NoteInfo) (int64, error) {
	return s.noteRepository.MultiAdd(ctx, notes)
}
