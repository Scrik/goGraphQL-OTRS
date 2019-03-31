package api

import "time"

type TypeUsers struct {
	Id         int32      `db:"id"`
	Login      *string    `db:"login"`
	Pw         *string    `db:"pw"`
	Title      *string    `db:"title"`
	FirstName  *string    `db:"first_name"`
	LastName   *string    `db:"last_name"`
	ValidId    *int8      `db:"valid_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
