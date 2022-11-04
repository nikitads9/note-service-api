package convert

import (
	"database/sql"

	"github.com/nikitads9/note-service-api/internal/app/model"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNoteInfo(note *desc.Note) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   note.GetTitle(),
		Content: note.GetContent(),
	}
}

func ToNotesInfo(notes []*desc.Note) []*model.NoteInfo {
	res := make([]*model.NoteInfo, 0, len(notes))

	for _, elem := range notes {
		res = append(res, ToNoteInfo(elem))
	}

	return res
}

func ToDescNoteInfo(note *model.NoteInfo) *desc.NoteInfo {
	res := &desc.NoteInfo{
		Id: note.Id,
		Note: &desc.Note{
			Title:   note.Title,
			Content: note.Content,
		},

		CreatedAt: timestamppb.New(note.CreatedAt),
	}

	if note.UpdatedAt.Valid {
		res.UpdatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return res
}

func ToDescNotesInfo(notes []*model.NoteInfo) []*desc.NoteInfo {
	res := make([]*desc.NoteInfo, 0, len(notes))
	for _, elem := range notes {
		res = append(res, ToDescNoteInfo(elem))
	}

	return res
}

func ToUpdateNoteInfo(req *desc.UpdateNoteRequest) *model.UpdateNoteInfo {
	return &model.UpdateNoteInfo{
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
