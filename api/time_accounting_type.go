package api

import "time"

type TypeTimeAccounting struct {
	Id         int64      `db:"id"`
	TicketId   *int64     `db:"ticket_id"`
	ArticleId  *int64     `db:"article_id"`
	TimeUnit   *string    `db:"time_unit"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
