package api

import "time"

type TypeProcessId struct {
	ProcessName   string  `db:"process_name"`
	ProcessId     *string `db:"process_id"`
	ProcessHost   *string `db:"process_host"`
	ProcessCreate *int32  `db:"process_create"`
	ProcessChange *int32  `db:"process_change"`
}
