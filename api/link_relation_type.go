package api

import "time"

type TypeLinkRelation struct {
	SourceObjectId int8       `db:"source_object_id"`
	SourceKey      *string    `db:"source_key"`
	TargetObjectId *int8      `db:"target_object_id"`
	TargetKey      *string    `db:"target_key"`
	TypeId         *int8      `db:"type_id"`
	StateId        *int8      `db:"state_id"`
	CreateTime     *time.Time `db:"create_time"`
	CreateBy       *int32     `db:"create_by"`
}
