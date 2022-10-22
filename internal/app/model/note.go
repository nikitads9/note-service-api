package model

import "database/sql"

type NoteInfo struct {
	Id      int64          `db:"id"`
	Title   sql.NullString `db:"title"`
	Content sql.NullString `db:"content"`
}
