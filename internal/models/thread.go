package models

import (
	"database/sql"
)

const (
	ThreadIdKey      = "id"
	ThreadAuthorKey  = "author"
	ThreadTitleKey   = "title"
	ThreadBodyKey    = "body"
	ThreadCreatedKey = "time_created"
	ThreadEditedKey  = "time_edited"
	ThreadTagsKey    = "tags"
)

type Thread struct {
	Id      int            `json:"id"`
	Author  sql.NullString `json:"author"`
	Title   string         `json:"title"`
	Body    string         `json:"body"`
	Created string         `json:"time_created"`
	Edited  sql.NullString `json:"time_edited"`
	Tags    string         `json:"tags"`
}
