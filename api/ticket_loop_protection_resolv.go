package api

type ResolverTicketLoopProtection struct {
	s TypeTicketLoopProtection
}

func (R *ResolverTicketLoopProtection) Set(s TypeTicketLoopProtection) {
	R.s = s
}

func (R ResolverTicketLoopProtection) SentTo() string {
	return R.s.SentTo
}

func (R ResolverTicketLoopProtection) SentDate() *string {
	return R.s.SentDate
}

func GenGqlTypeTicketLoopProtection(extra string) string {
	return "type TicketLoopProtection { " + extra + `
	SentTo: String,
	SentDate: String,
	}`
}
