package api

import "time"

type TypeNotificationEvent struct {
	Id         int32      `db:"id"`
	Name       *string    `db:"name"`
	ValidId    *int8      `db:"valid_id"`
	Comments   *string    `db:"comments"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
