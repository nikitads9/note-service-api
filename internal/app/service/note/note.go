package note

import "github.com/nikitads9/note-service-api/internal/app/repository/note_db"

type Service struct {
	noteRepository note_db.Repository
}

func NewNoteService(noteRepository note_db.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewMockNoteService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case note_db.Repository:
			is.noteRepository = s
		}
	}
	return &is
}
