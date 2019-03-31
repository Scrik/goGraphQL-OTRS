package api

type ResolverTicketWatcher struct {
	s TypeTicketWatcher
}

func (R *ResolverTicketWatcher) Set(s TypeTicketWatcher) {
	R.s = s
}

func (R ResolverTicketWatcher) TicketId() string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R ResolverTicketWatcher) UserId() *int32 {
	return R.s.UserId
}

func (R *ResolverTicketWatcher) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketWatcher) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketWatcher) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketWatcher) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketWatcher(extra string) string {
	return "type TicketWatcher { " + extra + `
	TicketId: String,
	UserId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
