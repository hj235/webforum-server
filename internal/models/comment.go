package models

import (
	"database/sql"
)

const (
	CommentIdKey       = "id"
	CommentNameKey     = "author"
	CommentThreadIdKey = "thread_id"
	CommentBodyKey     = "body"
	CommentCreatedKey  = "time_created"
	CommentEditedKey   = "time_edited"
)

type Comment struct {
	Id       int            `json:"id"`
	Author   sql.NullString `json:"author"`
	ThreadId int            `json:"thread_id"`
	Body     string         `json:"body"`
	Created  string         `json:"time_created"`
	Edited   sql.NullString `json:"time_edited"`
}
