package tiket

import (
	"fmt"
	"strings"
	"time"

	"github.com/goGraphQL-OTRS/api/tiket"
	"github.com/goGraphQL-OTRS/api/types"
	"github.com/goGraphQL-OTRS/internal/db"
)

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

func One(ID string) (result *tiket.Resolver, err error) {
	result = &tiket.Resolver{}
	fields := strings.Join([]string{
		"id", "tn", "title", "queue_id", "ticket_lock_id", "type_id", "service_id", "sla_id", "user_id", "responsible_user_id", "ticket_priority_id", "ticket_state_id", "customer_id", "customer_user_id", "timeout", "until_time", "escalation_time", "escalation_update_time", "escalation_response_time", "escalation_solution_time", "archive_flag", "create_time_unix", "create_by", "change_by",
	}, ", ")
	Tiket := types.TypeTiket{}
	row := db.DB.QueryRowx("select "+fields+" from `otrs`.`ticket` where id=?", args.ID)
	err = row.StructScan(&Tiket)
	if err != nil {
		return nil, err
	}
	result.Set(Tiket)
	return
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

func (R *Resolver) TypeId() *int32 {
	i := int32(*R.s.TypeId)
	return &i
}

func (R *Resolver) ServiceId() *int32 {
	return R.s.ServiceId
}

func (R *Resolver) SlaId() *int32 {
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

func (R *Resolver) CustomerId() *string {
	return R.s.CustomerId
}

func (R *Resolver) CustomerUserId() *string {
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

func (R *Resolver) CreateBy() int32 {
	return R.s.CreateBy
}

// func (R *Resolver) ChangeTime() *string {
// 	// return R.s.ChangeTime.String()
// 	s := "test"
// 	return &s
// }

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
