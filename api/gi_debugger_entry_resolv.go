package api

type ResolverGiDebuggerEntry struct {
	s TypeGiDebuggerEntry
}

func (R *ResolverGiDebuggerEntry) Set(s TypeGiDebuggerEntry) {
	R.s = s
}

func (R ResolverGiDebuggerEntry) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverGiDebuggerEntry) CommunicationId() *string {
	return R.s.CommunicationId
}

func (R ResolverGiDebuggerEntry) CommunicationType() *string {
	return R.s.CommunicationType
}

func (R ResolverGiDebuggerEntry) RemoteIp() *string {
	return R.s.RemoteIp
}

func (R ResolverGiDebuggerEntry) WebserviceId() *int32 {
	return R.s.WebserviceId
}

func (R *ResolverGiDebuggerEntry) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func GenGqlTypeGiDebuggerEntry(extra string) string {
	return "type GiDebuggerEntry { " + extra + `
	Id: String,
	CommunicationId: String,
	CommunicationType: String,
	RemoteIp: String,
	WebserviceId: Int,
	CreateTime: String,
	}`
}
