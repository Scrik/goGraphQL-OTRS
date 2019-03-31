package api

type ResolverTicketFlag struct {
	s TypeTicketFlag
}

func (R *ResolverTicketFlag) Set(s TypeTicketFlag) {
	R.s = s
}

func (R ResolverTicketFlag) TicketId() string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R ResolverTicketFlag) TicketKey() *string {
	return R.s.TicketKey
}

func (R ResolverTicketFlag) TicketValue() *string {
	return R.s.TicketValue
}

func (R *ResolverTicketFlag) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketFlag) CreateBy() *int32 {
	return R.s.CreateBy
}

func GenGqlTypeTicketFlag(extra string) string {
	return "type TicketFlag { " + extra + `
	TicketId: String,
	TicketKey: String,
	TicketValue: String,
	CreateTime: String,
	CreateBy: Int,
	}`
}
