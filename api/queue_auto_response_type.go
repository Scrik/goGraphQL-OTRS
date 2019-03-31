package api

import "time"

type TypeQueueAutoResponse struct {
	Id             int32      `db:"id"`
	QueueId        *int32     `db:"queue_id"`
	AutoResponseId *int32     `db:"auto_response_id"`
	CreateTime     *time.Time `db:"create_time"`
	CreateBy       *int32     `db:"create_by"`
	ChangeTime     *time.Time `db:"change_time"`
	ChangeBy       *int32     `db:"change_by"`
}
