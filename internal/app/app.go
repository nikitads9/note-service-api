package app

import (
	"context"
	"fmt"
	"log"

	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nikitads9/note-service-api/internal/app/api/note_v1"
	"github.com/nikitads9/note-service-api/internal/config"
	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	noteImpl        *note_v1.Implementation
	serviceProvider *serviceProvider
	pathConfig      string
	grpcServer      *grpc.Server
	mux             *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)
	return a, err
}

func (a *App) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	var err error
	cfg := a.serviceProvider.GetConfig()

	defer a.serviceProvider.db.Close()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = a.startGRPC(cfg)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = a.startHTTP(cfg)
	}()

	wg.Wait()

	return err

}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	noteService, err := a.serviceProvider.GetNoteService(ctx)
	if err != nil {
		return err
	}

	a.noteImpl = note_v1.NewNoteV1(noteService)

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))

	pb.RegisterNoteV1Server(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().Grpc.Port, opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) startGRPC(cfg *config.Config) error {
	list, err := net.Listen("tcp", cfg.Grpc.Port)
	if err != nil {
		return fmt.Errorf("failed to create listener %v", err.Error())
	}

	defer list.Close()

	if err = a.grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to process gRPC server: %s", err.Error())
	}

	return nil
}

func (a *App) startHTTP(cfg *config.Config) error {
	return http.ListenAndServe(cfg.Http.Port, a.mux)
}
