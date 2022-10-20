package convert

import (
	"database/sql"

	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToNoteInfo(req *desc.AddNoteRequest) *model.NoteInfo {
	return &model.NoteInfo{
		Title: sql.NullString{
			String: req.Note.GetTitle(),
			Valid:  true,
		},
		Content: sql.NullString{
			String: req.Note.GetContent(),
			Valid:  true,
		},
	}
}
