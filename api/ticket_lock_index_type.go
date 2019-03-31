package api

import "time"

type TypeTicketLockIndex struct {
	TicketId int64 `db:"ticket_id"`
}
