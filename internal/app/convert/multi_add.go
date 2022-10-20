package convert

import (
	"database/sql"

	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func ToNotesInfo(req *desc.MultiAddRequest) []*model.NoteInfo {
	notes := req.GetNotes()
	res := make([]*model.NoteInfo, 0, len(notes))

	for _, elem := range notes {
		res = append(res, &model.NoteInfo{
			Title: sql.NullString{
				String: elem.GetTitle(),
				Valid:  true,
			},
			Content: sql.NullString{
				String: elem.GetContent(),
				Valid:  true,
			},
		})
	}

	return res
}
