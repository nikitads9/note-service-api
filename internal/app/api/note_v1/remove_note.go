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

func (i *Implementation) RemoveNote(ctx context.Context, req *desc.RemoveNoteRequest) (*desc.Empty, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		log.Printf("failed to open connection to database %v\n", err.Error())
		return nil, err
	}
	defer db.Close()

	builder := sq.Delete(notesTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
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

	row.Next()
	var deleted int64
	err = row.Scan(&deleted)
	if err != nil {
		log.Printf("failed to get id of deleted entry %v\n", err.Error())
		return nil, err
	}

	if deleted == int64(0) {
		log.Printf("no entries removed\n")
		return &desc.Empty{}, nil
	}

	fmt.Printf("note with id=%v removed.\n", deleted)

	return &desc.Empty{}, nil
}
