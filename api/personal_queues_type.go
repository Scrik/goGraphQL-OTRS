package api

import "time"

type TypePersonalQueues struct {
	UserId  int32  `db:"user_id"`
	QueueId *int32 `db:"queue_id"`
}
