package app

import (
	"context"
	"fmt"
	"log"

	"github.com/nikitads9/note-service-api/internal/app/repository"
	"github.com/nikitads9/note-service-api/internal/app/service/note"
	"github.com/nikitads9/note-service-api/internal/config"
	"github.com/nikitads9/note-service-api/internal/pkg/db"
)

type serviceProvider struct {
	db             db.Client
	configPath     string
	config         *config.Config
	noteRepository repository.INoteRepository
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

func (s *serviceProvider) GetNoteRepository(ctx context.Context) (repository.INoteRepository, error) {
	if s.noteRepository == nil {
		s.noteRepository = repository.NewNoteRepository(s.GetDB(ctx))
		return s.noteRepository, nil
	}

	return s.noteRepository, nil
}

func (s *serviceProvider) GetNoteService(ctx context.Context) (*note.Service, error) {
	if s.noteService == nil {
		noteRepository, err := s.GetNoteRepository(ctx)
		if err != nil {
			return nil, fmt.Errorf("could note create repository err:%s", err)
		}
		s.noteService = note.NewNoteService(noteRepository)
	}

	return s.noteService, nil
}
