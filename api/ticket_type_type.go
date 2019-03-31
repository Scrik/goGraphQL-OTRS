package api

import "time"

type TypeTicketType struct {
	Id         int8       `db:"id"`
	Name       *string    `db:"name"`
	ValidId    *int8      `db:"valid_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
