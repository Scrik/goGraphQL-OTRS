package api

import "time"

type TypePmProcess struct {
	Id            int32      `db:"id"`
	EntityId      *string    `db:"entity_id"`
	Name          *string    `db:"name"`
	StateEntityId *string    `db:"state_entity_id"`
	Layout        *[]byte    `db:"layout"`
	Config        *[]byte    `db:"config"`
	CreateTime    *time.Time `db:"create_time"`
	CreateBy      *int32     `db:"create_by"`
	ChangeTime    *time.Time `db:"change_time"`
	ChangeBy      *int32     `db:"change_by"`
}
