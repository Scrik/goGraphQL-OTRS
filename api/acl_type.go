package api

import "time"

type TypeAcl struct {
	Id             int32      `db:"id"`
	Name           *string    `db:"name"`
	Comments       *string    `db:"comments"`
	Description    *string    `db:"description"`
	ValidId        *int8      `db:"valid_id"`
	StopAfterMatch *int8      `db:"stop_after_match"`
	ConfigMatch    *[]byte    `db:"config_match"`
	ConfigChange   *[]byte    `db:"config_change"`
	CreateTime     *time.Time `db:"create_time"`
	CreateBy       *int32     `db:"create_by"`
	ChangeTime     *time.Time `db:"change_time"`
	ChangeBy       *int32     `db:"change_by"`
}
