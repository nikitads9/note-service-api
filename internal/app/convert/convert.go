package convert

import (
	"database/sql"

	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNoteInfo(note *desc.Notes) *model.NoteInfo {
	return &model.NoteInfo{
		Title: sql.NullString{
			String: note.Title,
			Valid:  true,
		},
		Content: sql.NullString{
			String: note.Content,
			Valid:  true,
		},
	}
}

func ToNotesInfo(notes []*desc.Notes) []*model.NoteInfo {
	res := make([]*model.NoteInfo, 0, len(notes))

	for _, elem := range notes {
		res = append(res, ToNoteInfo(elem))
	}

	return res
}

func ToDescNoteInfo(note *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Id: note.Id,
		Note: &desc.Notes{
			Title:   note.Title.String,
			Content: note.Content.String,
		},
	}
}

func ToDescNotesInfo(notes []*model.NoteInfo) []*desc.NoteInfo {
	res := make([]*desc.NoteInfo, 0, len(notes))
	for _, elem := range notes {
		res = append(res, ToDescNoteInfo(elem))
	}

	return res
}

func ToUpdateNoteInfo(req *desc.UpdateNoteRequest) *model.NoteInfo {
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
