package note_v1

import (
	"github.com/nikitads9/note-service-api/internal/app/service/note"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService *note.Service
}

func NewNoteV1(noteService *note.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteV1Server{},
		noteService,
	}
}

func newMockNoteV1(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteV1Server{},
		i.noteService,
	}
}
