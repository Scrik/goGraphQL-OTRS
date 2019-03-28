package api

import (
	"fmt"
	"time"
)

type TypeTiket struct {
	Id                     int64     `json:"id"`
	Tn                     string    `json:"tn"`
	Title                  string    `json:"title"`
	QueueId                int32     `json:"queue_id"`
	TicketLockId           int8      `json:"ticket_lock_id"`
	TypeId                 int8      `json:"type_id"`
	ServiceId              int32     `json:"service_id"`
	SlaId                  int32     `json:"sla_id"`
	UserId                 int32     `json:"user_id"`
	ResponsibleUserId      int32     `json:"responsible_user_id"`
	TicketPriorityId       int8      `json:"ticket_priority_id"`
	TicketStateId          int8      `json:"ticket_state_id"`
	CustomerId             string    `json:"customer_id"`
	CustomerUserId         string    `json:"customer_user_id"`
	Timeout                int32     `json:"timeout"`
	UntilTime              int32     `json:"until_time"`
	EscalationTime         int32     `json:"escalation_time"`
	EscalationUpdateTime   int32     `json:"escalation_update_time"`
	EscalationResponseTime int32     `json:"escalation_response_time"`
	EscalationSolutionTime int32     `json:"escalation_solution_time"`
	ArchiveFlag            int8      `json:"archive_flag"`
	CreateTimeUnix         int64     `json:"create_time_unix"`
	CreateTime             time.Time `json:"create_time"`
	CreateBy               int32     `json:"create_by"`
	ChangeTime             time.Time `json:"change_time"`
	ChangeBy               int32     `json:"change_by"`
}

type Resolver struct {
	s TypeTiket
}

func (R *Resolver) Set(s TypeTiket) {
	R.s = s
}

func (R *Resolver) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *Resolver) Tn() string {
	return R.s.Tn
}

func (R *Resolver) Title() string {
	return R.s.Title
}

func (R *Resolver) QueueId() int32 {
	return R.s.QueueId
}

func (R *Resolver) TicketLockId() int32 {
	return int32(R.s.TicketLockId)
}

func (R *Resolver) TypeId() int32 {
	return int32(R.s.TypeId)
}

func (R *Resolver) ServiceId() int32 {
	return R.s.ServiceId
}

func (R *Resolver) SlaId() int32 {
	return R.s.SlaId
}

func (R *Resolver) UserId() int32 {
	return R.s.UserId
}

func (R *Resolver) ResponsibleUserId() int32 {
	return R.s.ResponsibleUserId
}

func (R *Resolver) TicketPriorityId() int32 {
	return int32(R.s.TicketPriorityId)
}

func (R *Resolver) TicketStateId() int32 {
	return int32(R.s.TicketStateId)
}

func (R *Resolver) CustomerId() string {
	return R.s.CustomerId
}

func (R *Resolver) CustomerUserId() string {
	return R.s.CustomerUserId
}

func (R *Resolver) Timeout() int32 {
	return R.s.Timeout
}

func (R *Resolver) UntilTime() int32 {
	return R.s.UntilTime
}

func (R *Resolver) EscalationTime() int32 {
	return R.s.EscalationTime
}

func (R *Resolver) EscalationUpdateTime() int32 {
	return R.s.EscalationUpdateTime
}

func (R *Resolver) EscalationResponseTime() int32 {
	return R.s.EscalationResponseTime
}

func (R *Resolver) EscalationSolutionTime() int32 {
	return R.s.EscalationSolutionTime
}

func (R *Resolver) ArchiveFlag() int32 {
	return int32(R.s.ArchiveFlag)
}

func (R *Resolver) CreateTimeUnix() string {
	return fmt.Sprintf("%d", R.s.CreateTimeUnix)
}

func (R *Resolver) CreateTime() string {
	return R.s.CreateTime.String()
}

func (R *Resolver) CreateBy() int32 {
	return R.s.CreateBy
}

func (R *Resolver) ChangeTime() string {
	return R.s.ChangeTime.String()
}

func (R *Resolver) ChangeBy() int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTiket(extra string) string {
	return "type Tiket { " + extra + `
	Id: String!,
	Tn: String!,
	Title: String!,
	QueueId: Int!,
	TicketLockId: Int!,
	TypeId: Int!,
	ServiceId: Int!,
	SlaId: Int!,
	UserId: Int!,
	ResponsibleUserId: Int!,
	TicketPriorityId: Int!,
	TicketStateId: Int!,
	CustomerId: String!,
	CustomerUserId: String!,
	Timeout: Int!,
	UntilTime: Int!,
	EscalationTime: Int!,
	EscalationUpdateTime: Int!,
	EscalationResponseTime: Int!,
	EscalationSolutionTime: Int!,
	ArchiveFlag: Int!,
	CreateTimeUnix: String!,
	CreateTime: String!,
	CreateBy: Int!,
	ChangeTime: String!,
	ChangeBy: Int!,
	}`
}
