package api

import "time"

type TypeVirtualFs struct {
	Id         int64      `db:"id"`
	Filename   *string    `db:"filename"`
	Backend    *string    `db:"backend"`
	BackendKey *string    `db:"backend_key"`
	CreateTime *time.Time `db:"create_time"`
}
