package api

import "time"

type TypeQueue struct {
	Id                  int32      `db:"id"`
	Name                *string    `db:"name"`
	GroupId             *int32     `db:"group_id"`
	UnlockTimeout       *int32     `db:"unlock_timeout"`
	FirstResponseTime   *int32     `db:"first_response_time"`
	FirstResponseNotify *int8      `db:"first_response_notify"`
	UpdateTime          *int32     `db:"update_time"`
	UpdateNotify        *int8      `db:"update_notify"`
	SolutionTime        *int32     `db:"solution_time"`
	SolutionNotify      *int8      `db:"solution_notify"`
	SystemAddressId     *int8      `db:"system_address_id"`
	CalendarName        *string    `db:"calendar_name"`
	DefaultSignKey      *string    `db:"default_sign_key"`
	SalutationId        *int8      `db:"salutation_id"`
	SignatureId         *int8      `db:"signature_id"`
	FollowUpId          *int8      `db:"follow_up_id"`
	FollowUpLock        *int8      `db:"follow_up_lock"`
	Comments            *string    `db:"comments"`
	ValidId             *int8      `db:"valid_id"`
	CreateTime          *time.Time `db:"create_time"`
	CreateBy            *int32     `db:"create_by"`
	ChangeTime          *time.Time `db:"change_time"`
	ChangeBy            *int32     `db:"change_by"`
}
