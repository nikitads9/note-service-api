package note_repository

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/note_service_repository.go -package=mocks . Repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/nikitads9/note-service-api/internal/app/model"
	"github.com/nikitads9/note-service-api/internal/app/repository/table"
	"github.com/nikitads9/note-service-api/internal/pkg/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "there is no note with this id")

type Repository interface {
	AddNote(ctx context.Context, note *model.NoteInfo) (int64, error)
	GetList(ctx context.Context) ([]*model.NoteInfo, error)
	GetNote(ctx context.Context, id int64) (*model.NoteInfo, error)
	MultiAdd(ctx context.Context, notes []*model.NoteInfo) (int64, error)
	RemoveNote(ctx context.Context, id int64) (int64, error)
	UpdateNote(ctx context.Context, note *model.UpdateNoteInfo) error
}
type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) AddNote(ctx context.Context, note *model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.NotesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content, created_at").
		Values(note.Title, note.Content, time.Now().UTC()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "AddNoteRepo",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
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

func (r *repository) GetList(ctx context.Context) ([]*model.NoteInfo, error) {
	builder := sq.Select("id, title, content, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.NotesTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetListRepo",
		QueryRaw: query,
	}

	var res []*model.NoteInfo
	err = r.client.DB().SelectContext(ctx, &res, q, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetNote(ctx context.Context, id int64) (*model.NoteInfo, error) {
	builder := sq.Select("id, title, content, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.NotesTable).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNoteRepo",
		QueryRaw: query,
	}

	var res = new(model.NoteInfo)
	err = r.client.DB().GetContext(ctx, res, q, args...)
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return res, nil
}

func (r *repository) MultiAdd(ctx context.Context, notes []*model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.NotesTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content, created_at").
		Suffix("returning id")

	for _, note := range notes {
		builder = builder.Values(note.Title, note.Content, time.Now().UTC())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "MultiAddRepo",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
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

func (r *repository) RemoveNote(ctx context.Context, id int64) (int64, error) {
	builder := sq.Delete(table.NotesTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "RemoveNoteRepo",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	isNotEmpty := row.Next()
	var removedID int64
	err = row.Scan(&removedID)
	if err != nil {
		if !isNotEmpty {
			return 0, errNotFound
		}
		return 0, err
	}

	return removedID, nil
}

func (r *repository) UpdateNote(ctx context.Context, note *model.UpdateNoteInfo) error {
	builder := sq.Update(table.NotesTable).
		Set("updated_at", time.Now().UTC()).
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

	q := db.Query{
		Name:     "UpdateNoteRepo",
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errNotFound
	}

	return nil
}
