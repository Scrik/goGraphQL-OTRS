package api

import "time"

type TypeGiObjectLockState struct {
	WebserviceId     int32      `db:"webservice_id"`
	ObjectType       *string    `db:"object_type"`
	ObjectId         *int64     `db:"object_id"`
	LockState        *string    `db:"lock_state"`
	LockStateCounter *int32     `db:"lock_state_counter"`
	CreateTime       *time.Time `db:"create_time"`
	ChangeTime       *time.Time `db:"change_time"`
}
