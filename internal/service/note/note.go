package note

import "github.com/nikitads9/note-service-api/internal/repository/note"

type Service struct {
	noteRepository note.Repository
}

func NewNoteService(noteRepository note.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewMockNoteService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case note.Repository:
			is.noteRepository = s
		}
	}
	return &is
}
