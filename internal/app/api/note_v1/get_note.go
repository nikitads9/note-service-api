package note_v1

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		log.Printf("failed to open connection to database %v\n", err.Error())
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id, title, content").
		PlaceholderFormat(sq.Dollar).
		From("notes").
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("failed to build a query %v\n", err.Error())
		return nil, err
	}

	type note struct {
		Id      int64  `db:"id"`
		Title   string `db:"title"`
		Content string `db:"content"`
	}
	var res []note

	err = db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		log.Printf("failed to select %v\n", err.Error())
		return nil, err
	}

	if len(res) == 0 {
		return nil, status.Error(codes.NotFound, "ebanyi rot etogo kasino")
	}

	fmt.Printf("requested note with id %v\n", req.Id)

	return &desc.GetNoteResponse{
		Id:      res[0].Id,
		Title:   res[0].Title,
		Content: res[0].Content,
	}, nil
}
