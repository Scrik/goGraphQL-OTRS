package api

import "time"

type TypePmEntitySync struct {
	EntityType string     `db:"entity_type"`
	EntityId   *string    `db:"entity_id"`
	SyncState  *string    `db:"sync_state"`
	CreateTime *time.Time `db:"create_time"`
	ChangeTime *time.Time `db:"change_time"`
}
