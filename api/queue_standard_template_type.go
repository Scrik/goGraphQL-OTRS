package api

import "time"

type TypeQueueStandardTemplate struct {
	QueueId            int32      `db:"queue_id"`
	StandardTemplateId *int32     `db:"standard_template_id"`
	CreateTime         *time.Time `db:"create_time"`
	CreateBy           *int32     `db:"create_by"`
	ChangeTime         *time.Time `db:"change_time"`
	ChangeBy           *int32     `db:"change_by"`
}
