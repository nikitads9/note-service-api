package app

import (
	"context"
	"log"

	"github.com/nikitads9/note-service-api/internal/app/repository/note_db"
	"github.com/nikitads9/note-service-api/internal/app/service/note"
	"github.com/nikitads9/note-service-api/internal/config"
	"github.com/nikitads9/note-service-api/internal/pkg/db"
)

type serviceProvider struct {
	db             db.Client
	configPath     string
	config         *config.Config
	noteRepository note_db.Repository
	noteService    *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("could not get config err: %s", err.Error())
		}
		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can`t connect to db err: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.Read(s.configPath)
		if err != nil {
			log.Fatalf("could not get config err: %s", err)
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) note_db.Repository {
	if s.noteRepository == nil {
		s.noteRepository = note_db.NewNoteRepository(s.GetDB(ctx))
		return s.noteRepository
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		noteRepository := s.GetNoteRepository(ctx)
		s.noteService = note.NewNoteService(noteRepository)
	}

	return s.noteService
}
