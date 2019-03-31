package api

import "time"

type TypeWebUploadCache struct {
	FormId         string  `db:"form_id"`
	Filename       *string `db:"filename"`
	ContentId      *string `db:"content_id"`
	ContentSize    *string `db:"content_size"`
	ContentType    *string `db:"content_type"`
	Disposition    *string `db:"disposition"`
	Content        *[]byte `db:"content"`
	CreateTimeUnix *int64  `db:"create_time_unix"`
}
