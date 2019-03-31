package api

import "time"

type TypeTicketWatcher struct {
	TicketId   int64      `db:"ticket_id"`
	UserId     *int32     `db:"user_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
