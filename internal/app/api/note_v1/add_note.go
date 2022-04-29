package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

const (
	notesTable = "notes"
	host       = "localhost"
	port       = "5444"
	user       = "postgres"
	password   = "notes_pass"
	dbName     = "notes_db"
	ssl        = "disable"
)

func (i *Implementation) AddNote(ctx context.Context, req *desc.AddNoteRequest) (*desc.AddNoteResponse, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Insert(notesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content").
		Values(req.GetTitle(), req.GetContent()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.AddNoteResponse{
		Result: &desc.AddNoteResponse_Result{
			Id: id,
		},
	}, nil
}
