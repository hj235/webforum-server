package models

const (
	TagThreadIdKey = "thread_id"
	TagBodyKey     = "body"
)

type Tag struct {
	ThreadId int    `json:"thread_id"`
	Body     string `json:"body"`
}
