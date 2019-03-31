package api

type ResolverTicketIndex struct {
	s TypeTicketIndex
}

func (R *ResolverTicketIndex) Set(s TypeTicketIndex) {
	R.s = s
}

func (R ResolverTicketIndex) TicketId() string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R ResolverTicketIndex) QueueId() *int32 {
	return R.s.QueueId
}

func (R ResolverTicketIndex) Queue() *string {
	return R.s.Queue
}

func (R ResolverTicketIndex) GroupId() *int32 {
	return R.s.GroupId
}

func (R ResolverTicketIndex) SLock() *string {
	return R.s.SLock
}

func (R ResolverTicketIndex) SState() *string {
	return R.s.SState
}

func (R *ResolverTicketIndex) CreateTimeUnix() *string {
	str := fmt.Sprintf("%d", R.s.CreateTimeUnix)
	return &str
}

func GenGqlTypeTicketIndex(extra string) string {
	return "type TicketIndex { " + extra + `
	TicketId: String,
	QueueId: Int,
	Queue: String,
	GroupId: Int,
	SLock: String,
	SState: String,
	CreateTimeUnix: String,
	}`
}
