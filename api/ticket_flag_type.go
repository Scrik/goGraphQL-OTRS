package api

import "time"

type TypeTicketFlag struct {
	TicketId    int64      `db:"ticket_id"`
	TicketKey   *string    `db:"ticket_key"`
	TicketValue *string    `db:"ticket_value"`
	CreateTime  *time.Time `db:"create_time"`
	CreateBy    *int32     `db:"create_by"`
}
