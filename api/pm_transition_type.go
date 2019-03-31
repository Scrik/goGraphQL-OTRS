package api

import "time"

type TypePmTransition struct {
	Id         int32      `db:"id"`
	EntityId   *string    `db:"entity_id"`
	Name       *string    `db:"name"`
	Config     *[]byte    `db:"config"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
