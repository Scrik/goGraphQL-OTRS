package api

import "time"

type TypeTicketState struct {
	Id         int8       `db:"id"`
	Name       *string    `db:"name"`
	Comments   *string    `db:"comments"`
	TypeId     *int8      `db:"type_id"`
	ValidId    *int8      `db:"valid_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
