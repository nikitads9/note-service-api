package note_v1

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		log.Printf("failed to open connection to database %v\n", err.Error())
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(notesTable).
		Set("title", req.GetTitle()).
		Set("content", req.GetContent()).
		Where(sq.Eq{"id": req.GetId()}).
		PlaceholderFormat(sq.Dollar).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("failed to build a query %v\n", err.Error())
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("failed to get query context %v\n", err.Error())
		return nil, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		log.Printf("failed get updated entry id %v\n", err.Error())
		return nil, err
	}

	fmt.Printf("edited note with id %v\n", id)

	return &desc.Empty{}, nil
}
