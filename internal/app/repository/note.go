package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/note_service_repository.go -package=mocks . INoteRepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/nikitads9/note-service-api/internal/app/model"
	"github.com/nikitads9/note-service-api/internal/app/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type INoteRepository interface {
	AddNote(ctx context.Context, note *model.NoteInfo) (int64, error)
	GetList(ctx context.Context) ([]*model.NoteInfo, error)
	GetNote(ctx context.Context, id int64) (*model.NoteInfo, error)
	MultiAdd(ctx context.Context, notes []*model.NoteInfo) (int64, error)
	RemoveNote(ctx context.Context, id int64) (int64, error)
	UpdateNote(ctx context.Context, note *model.NoteInfo) error
}
type noteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) INoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (n *noteRepository) AddNote(ctx context.Context, note *model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.NotesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content").
		Values(note.Title.String, note.Content.String).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (n *noteRepository) GetList(ctx context.Context) ([]*model.NoteInfo, error) {
	builder := sq.Select("id, title, content").
		PlaceholderFormat(sq.Dollar).
		From(table.NotesTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []*model.NoteInfo

	err = n.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (n *noteRepository) GetNote(ctx context.Context, id int64) (*model.NoteInfo, error) {
	builder := sq.Select("id, title, content").
		PlaceholderFormat(sq.Dollar).
		From(table.NotesTable).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []*model.NoteInfo

	err = n.db.GetContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, status.Error(codes.NotFound, "ebanyi rot etogo kasino")
	}

	return res[0], nil
}

func (n *noteRepository) MultiAdd(ctx context.Context, notes []*model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.NotesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content").
		Suffix("returning id")

	for _, note := range notes {
		builder = builder.Values(note.Title.String, note.Content.String)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var added int64
	for row.Next() {
		added += 1
	}

	return added, nil
}

func (n *noteRepository) RemoveNote(ctx context.Context, id int64) (int64, error) {
	builder := sq.Delete(table.NotesTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var removedID int64
	err = row.Scan(&removedID)
	if err != nil {
		return 0, err
	}

	return removedID, nil
}

func (n *noteRepository) UpdateNote(ctx context.Context, note *model.NoteInfo) error {
	builder := sq.Update(table.NotesTable).
		Where(sq.Eq{"id": note.Id}).
		PlaceholderFormat(sq.Dollar).
		Suffix("returning id")
	if note.Title.Valid {
		builder = builder.Set("title", note.Title.String)
	}

	if note.Content.Valid {
		builder = builder.Set("content", note.Content.String)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
