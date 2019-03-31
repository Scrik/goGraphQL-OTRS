package api

import "time"

type TypeRoleUser struct {
	UserId     int32      `db:"user_id"`
	RoleId     *int32     `db:"role_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
