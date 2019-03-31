package api

import "time"

type TypeSchedulerRecurrentTask struct {
	Id                    int64      `db:"id"`
	Name                  *string    `db:"name"`
	TaskType              *string    `db:"task_type"`
	LastExecutionTime     *time.Time `db:"last_execution_time"`
	LastWorkerTaskId      *int64     `db:"last_worker_task_id"`
	LastWorkerStatus      *int8      `db:"last_worker_status"`
	LastWorkerRunningTime *int32     `db:"last_worker_running_time"`
	LockKey               *int64     `db:"lock_key"`
	LockTime              *time.Time `db:"lock_time"`
	CreateTime            *time.Time `db:"create_time"`
	ChangeTime            *time.Time `db:"change_time"`
}
