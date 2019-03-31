package api

type ResolverGenericAgentJobs struct {
	s TypeGenericAgentJobs
}

func (R *ResolverGenericAgentJobs) Set(s TypeGenericAgentJobs) {
	R.s = s
}

func (R ResolverGenericAgentJobs) JobName() string {
	return R.s.JobName
}

func (R ResolverGenericAgentJobs) JobKey() *string {
	return R.s.JobKey
}

func (R ResolverGenericAgentJobs) JobValue() *string {
	return R.s.JobValue
}

func GenGqlTypeGenericAgentJobs(extra string) string {
	return "type GenericAgentJobs { " + extra + `
	JobName: String,
	JobKey: String,
	JobValue: String,
	}`
}
