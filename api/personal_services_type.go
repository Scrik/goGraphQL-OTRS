package api

import "time"

type TypePersonalServices struct {
	UserId    int32  `db:"user_id"`
	ServiceId *int32 `db:"service_id"`
}
