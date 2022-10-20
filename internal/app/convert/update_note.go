package convert

import (
	"database/sql"

	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNoteUpdateInfo(req *desc.UpdateNoteRequest) *model.NoteInfo {
	return &model.NoteInfo{
		Id: req.GetId(),
		Title: sql.NullString{
			String: req.GetTitle().GetValue(),
			Valid:  req.GetTitle() != nil,
		},
		Content: sql.NullString{
			String: req.GetContent().GetValue(),
			Valid:  req.GetContent() != nil,
		},
	}
}
