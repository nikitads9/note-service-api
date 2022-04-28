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

type result struct {
	Title   string `db:"title"`
	Content string `db:"content"`
}

func (i *Implementation) GetList(ctx context.Context, in *desc.Empty) (*desc.GetListResponse, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		log.Printf("failed to open connection to database %v\n", err.Error())
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("title, content").
		PlaceholderFormat(sq.Dollar).
		From("notes")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("failed to build a query %v\n", err.Error())
		return nil, err
	}

	var res []result

	err = db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		log.Printf("failed to select %v\n", err.Error())
		return nil, err
	}

	notes := make([]*desc.GetListResponse_Result, len(res))
	for _, u := range res {
		notes = append(notes, &desc.GetListResponse_Result{
			Title:   u.Title,
			Content: u.Content,
		})
	}

	fmt.Println("returned all notes")

	return &desc.GetListResponse{
		Results: notes}, nil
}
