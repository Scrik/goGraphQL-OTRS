package api

import "time"

type TypeGroupRole struct {
	RoleId          int32      `db:"role_id"`
	GroupId         *int32     `db:"group_id"`
	PermissionKey   *string    `db:"permission_key"`
	PermissionValue *int8      `db:"permission_value"`
	CreateTime      *time.Time `db:"create_time"`
	CreateBy        *int32     `db:"create_by"`
	ChangeTime      *time.Time `db:"change_time"`
	ChangeBy        *int32     `db:"change_by"`
}
