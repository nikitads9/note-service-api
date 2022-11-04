package note

import "github.com/nikitads9/note-service-api/internal/app/repository/note_repository"

type Service struct {
	noteRepository note_repository.Repository
}

func NewNoteService(noteRepository note_repository.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewMockNoteService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case note_repository.Repository:
			is.noteRepository = s
		}
	}
	return &is
}
