package api

import "time"

type TypeAclSync struct {
	AclId      string     `db:"acl_id"`
	SyncState  *string    `db:"sync_state"`
	CreateTime *time.Time `db:"create_time"`
	ChangeTime *time.Time `db:"change_time"`
}
