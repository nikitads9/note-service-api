package note_v1

import (
	serv "github.com/nikitads9/note-service-api/internal/app/service/note"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService *serv.Service
}

func NewNoteV1(noteService *serv.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteV1Server{},
		noteService,
	}
}
