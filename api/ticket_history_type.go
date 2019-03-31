package api

import "time"

type TypeTicketHistory struct {
	Id            int64      `db:"id"`
	Name          *string    `db:"name"`
	HistoryTypeId *int8      `db:"history_type_id"`
	TicketId      *int64     `db:"ticket_id"`
	ArticleId     *int64     `db:"article_id"`
	TypeId        *int8      `db:"type_id"`
	QueueId       *int32     `db:"queue_id"`
	OwnerId       *int32     `db:"owner_id"`
	PriorityId    *int8      `db:"priority_id"`
	StateId       *int8      `db:"state_id"`
	CreateTime    *time.Time `db:"create_time"`
	CreateBy      *int32     `db:"create_by"`
	ChangeTime    *time.Time `db:"change_time"`
	ChangeBy      *int32     `db:"change_by"`
}
