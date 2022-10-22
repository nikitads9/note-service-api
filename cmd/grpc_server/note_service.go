package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/nikitads9/note-service-api/internal/app/api/note_v1"
	"github.com/nikitads9/note-service-api/internal/app/repository"
	"github.com/nikitads9/note-service-api/internal/app/service/note"
	"github.com/nikitads9/note-service-api/internal/config"
	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cfg *config.Config
	var wg sync.WaitGroup

	flag.Parse()

	cfg, err := config.Read("config.yml")
	if err != nil {
		log.Fatal("failed to open configuration file ", err)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startGRPC(cfg))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startHTTP(cfg))
	}()

	wg.Wait()
}

func startGRPC(cfg *config.Config) error {
	//nolint
	list, err := net.Listen("tcp", cfg.Grpc.Port)
	if err != nil {
		return fmt.Errorf("failed to create listener %v", err.Error())
	}
	defer list.Close()

	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		return fmt.Errorf("failed to establish connection with database")
	}
	defer db.Close()

	noteRepository := repository.NewNoteRepository(db)
	noteService := note.NewNoteService(noteRepository)

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	pb.RegisterNoteV1Server(s, note_v1.NewNoteV1(noteService))
	if err = s.Serve(list); err != nil {
		return fmt.Errorf("failed to process gRPC server %v", err.Error())
	}

	return nil
}

func startHTTP(cfg *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterNoteV1HandlerFromEndpoint(ctx, mux, cfg.Grpc.Port, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Http.Port, mux)
}
