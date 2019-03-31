package api

type ResolverProcessId struct {
	s TypeProcessId
}

func (R *ResolverProcessId) Set(s TypeProcessId) {
	R.s = s
}

func (R ResolverProcessId) ProcessName() string {
	return R.s.ProcessName
}

func (R ResolverProcessId) ProcessId() *string {
	return R.s.ProcessId
}

func (R ResolverProcessId) ProcessHost() *string {
	return R.s.ProcessHost
}

func (R ResolverProcessId) ProcessCreate() *int32 {
	return R.s.ProcessCreate
}

func (R ResolverProcessId) ProcessChange() *int32 {
	return R.s.ProcessChange
}

func GenGqlTypeProcessId(extra string) string {
	return "type ProcessId { " + extra + `
	ProcessName: String,
	ProcessId: String,
	ProcessHost: String,
	ProcessCreate: Int,
	ProcessChange: Int,
	}`
}
