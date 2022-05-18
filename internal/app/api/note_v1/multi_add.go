package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/nikitads9/note-service-api/pkg/note_api"
)

func (i *Implementation) MultiAdd(ctx context.Context, req *desc.MultiAddRequest) (*desc.MultiAddResponse, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Insert(notesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content").
		Suffix("returning id")

	for _, note := range req.GetNotes() {
		builder = builder.Values(note.GetTitle(), note.GetContent())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	added := []int64{}
	for row.Next() {
		var element int64
		err = row.Scan(&element)
		if err != nil {
			return nil, err
		}
		added = append(added, element)
	}

	return &desc.MultiAddResponse{
		Result: &desc.MultiAddResponse_Result{
			Count: int64(len(added)),
		},
	}, nil
}
