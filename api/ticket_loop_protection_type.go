package api

import "time"

type TypeTicketLoopProtection struct {
	SentTo   string  `db:"sent_to"`
	SentDate *string `db:"sent_date"`
}
