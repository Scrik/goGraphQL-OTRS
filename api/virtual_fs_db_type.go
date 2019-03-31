package api

import "time"

type TypeVirtualFsDb struct {
	Id         int64      `db:"id"`
	Filename   *string    `db:"filename"`
	Content    *[]byte    `db:"content"`
	CreateTime *time.Time `db:"create_time"`
}
