package model

import (
	"database/sql"
	"time"
)

type UpdateNoteInfo struct {
	Id      int64          `db:"id"`
	Title   sql.NullString `db:"title"`
	Content sql.NullString `db:"content"`
}

type NoteInfo struct {
	Id        int64     `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
