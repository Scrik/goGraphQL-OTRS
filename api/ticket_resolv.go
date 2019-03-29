package api

import "fmt"

type ResolverTicket struct {
	s TypeTicket
}

func (R *ResolverTicket) Set(s TypeTicket) {
	R.s = s
}

func (R ResolverTicket) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicket) Tn() *string {
	return R.s.Tn
}

func (R ResolverTicket) Title() *string {
	return R.s.Title
}

func (R ResolverTicket) QueueId() *int32 {
	return R.s.QueueId
}

func (R ResolverTicket) TicketLockId() *int32 {
	i := int32(*R.s.TicketLockId)
	return &i
}

func (R ResolverTicket) TypeId() *int32 {
	i := int32(*R.s.TypeId)
	return &i
}

func (R ResolverTicket) ServiceId() *int32 {
	return R.s.ServiceId
}

func (R ResolverTicket) SlaId() *int32 {
	return R.s.SlaId
}

func (R ResolverTicket) UserId() *int32 {
	return R.s.UserId
}

func (R ResolverTicket) ResponsibleUserId() *int32 {
	return R.s.ResponsibleUserId
}

func (R ResolverTicket) TicketPriorityId() *int32 {
	i := int32(*R.s.TicketPriorityId)
	return &i
}

func (R ResolverTicket) TicketStateId() *int32 {
	i := int32(*R.s.TicketStateId)
	return &i
}

func (R ResolverTicket) CustomerId() *string {
	return R.s.CustomerId
}

func (R ResolverTicket) CustomerUserId() *string {
	return R.s.CustomerUserId
}

func (R ResolverTicket) Timeout() *int32 {
	return R.s.Timeout
}

func (R ResolverTicket) UntilTime() *int32 {
	return R.s.UntilTime
}

func (R ResolverTicket) EscalationTime() *int32 {
	return R.s.EscalationTime
}

func (R ResolverTicket) EscalationUpdateTime() *int32 {
	return R.s.EscalationUpdateTime
}

func (R ResolverTicket) EscalationResponseTime() *int32 {
	return R.s.EscalationResponseTime
}

func (R ResolverTicket) EscalationSolutionTime() *int32 {
	return R.s.EscalationSolutionTime
}

func (R ResolverTicket) ArchiveFlag() *int32 {
	i := int32(*R.s.ArchiveFlag)
	return &i
}

func (R *ResolverTicket) CreateTimeUnix() *string {
	str := fmt.Sprintf("%d", R.s.CreateTimeUnix)
	return &str
}

func (R *ResolverTicket) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicket) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicket) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicket) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicket(extra string) string {
	return "type Ticket { " + extra + `
	Id: String,
	Tn: String,
	Title: String,
	QueueId: Int,
	TicketLockId: Int,
	TypeId: Int,
	ServiceId: Int,
	SlaId: Int,
	UserId: Int,
	ResponsibleUserId: Int,
	TicketPriorityId: Int,
	TicketStateId: Int,
	CustomerId: String,
	CustomerUserId: String,
	Timeout: Int,
	UntilTime: Int,
	EscalationTime: Int,
	EscalationUpdateTime: Int,
	EscalationResponseTime: Int,
	EscalationSolutionTime: Int,
	ArchiveFlag: Int,
	CreateTimeUnix: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
