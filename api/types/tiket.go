package types

import "time"

type TypeTiket struct {
	Id                     int64     `db:"id"`
	Tn                     string    `db:"tn"`
	Title                  string    `db:"title"`
	QueueId                int32     `db:"queue_id"`
	TicketLockId           int8      `db:"ticket_lock_id"`
	TypeId                 *int8     `db:"type_id"`
	ServiceId              *int32    `db:"service_id"`
	SlaId                  *int32    `db:"sla_id"`
	UserId                 int32     `db:"user_id"`
	ResponsibleUserId      int32     `db:"responsible_user_id"`
	TicketPriorityId       int8      `db:"ticket_priority_id"`
	TicketStateId          int8      `db:"ticket_state_id"`
	CustomerId             *string   `db:"customer_id"`
	CustomerUserId         *string   `db:"customer_user_id"`
	Timeout                int32     `db:"timeout"`
	UntilTime              int32     `db:"until_time"`
	EscalationTime         int32     `db:"escalation_time"`
	EscalationUpdateTime   int32     `db:"escalation_update_time"`
	EscalationResponseTime int32     `db:"escalation_response_time"`
	EscalationSolutionTime int32     `db:"escalation_solution_time"`
	ArchiveFlag            int8      `db:"archive_flag"`
	CreateTimeUnix         int64     `db:"create_time_unix"`
	CreateBy               int32     `db:"create_by"`
	ChangeBy               int32     `db:"change_by"`
	CreateTime             time.Time `db:"create_time"`
}
