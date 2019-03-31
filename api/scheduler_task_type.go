package api

import "time"

type TypeSchedulerTask struct {
	Id             int64      `db:"id"`
	Ident          *int64     `db:"ident"`
	Name           *string    `db:"name"`
	TaskType       *string    `db:"task_type"`
	TaskData       *[]byte    `db:"task_data"`
	Attempts       *int8      `db:"attempts"`
	LockKey        *int64     `db:"lock_key"`
	LockTime       *time.Time `db:"lock_time"`
	LockUpdateTime *time.Time `db:"lock_update_time"`
	CreateTime     *time.Time `db:"create_time"`
}
