package api

type ResolverTicketLockIndex struct {
	s TypeTicketLockIndex
}

func (R *ResolverTicketLockIndex) Set(s TypeTicketLockIndex) {
	R.s = s
}

func (R ResolverTicketLockIndex) TicketId() string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func GenGqlTypeTicketLockIndex(extra string) string {
	return "type TicketLockIndex { " + extra + `
	TicketId: String,
	}`
}
