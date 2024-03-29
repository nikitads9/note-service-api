package note

import (
	"context"

	"github.com/nikitads9/note-service-api/internal/model"
)

func (s *Service) UpdateNote(ctx context.Context, note *model.UpdateNoteInfo) error {
	return s.noteRepository.UpdateNote(ctx, note)
}
