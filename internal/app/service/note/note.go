package note

import "github.com/nikitads9/note-service-api/internal/app/repository"

type Service struct {
	noteRepository repository.INoteRepository
}

func NewNoteService(noteRepository repository.INoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewMockNoteService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case repository.INoteRepository:
			is.noteRepository = s
		}
	}
	return &is
}
