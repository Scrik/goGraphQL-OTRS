package api

import "time"

type TypeTicketIndex struct {
	TicketId       int64   `db:"ticket_id"`
	QueueId        *int32  `db:"queue_id"`
	Queue          *string `db:"queue"`
	GroupId        *int32  `db:"group_id"`
	SLock          *string `db:"s_lock"`
	SState         *string `db:"s_state"`
	CreateTimeUnix *int64  `db:"create_time_unix"`
}
