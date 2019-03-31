package api

type ResolverServiceSla struct {
	s TypeServiceSla
}

func (R *ResolverServiceSla) Set(s TypeServiceSla) {
	R.s = s
}

func (R ResolverServiceSla) ServiceId() int32 {
	return R.s.ServiceId
}

func (R ResolverServiceSla) SlaId() *int32 {
	return R.s.SlaId
}

func GenGqlTypeServiceSla(extra string) string {
	return "type ServiceSla { " + extra + `
	ServiceId: Int,
	SlaId: Int,
	}`
}
