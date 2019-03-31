package api

import "time"

type TypeSla struct {
	Id                  int32      `db:"id"`
	Name                *string    `db:"name"`
	CalendarName        *string    `db:"calendar_name"`
	FirstResponseTime   *int32     `db:"first_response_time"`
	FirstResponseNotify *int8      `db:"first_response_notify"`
	UpdateTime          *int32     `db:"update_time"`
	UpdateNotify        *int8      `db:"update_notify"`
	SolutionTime        *int32     `db:"solution_time"`
	SolutionNotify      *int8      `db:"solution_notify"`
	ValidId             *int8      `db:"valid_id"`
	Comments            *string    `db:"comments"`
	CreateTime          *time.Time `db:"create_time"`
	CreateBy            *int32     `db:"create_by"`
	ChangeTime          *time.Time `db:"change_time"`
	ChangeBy            *int32     `db:"change_by"`
}
