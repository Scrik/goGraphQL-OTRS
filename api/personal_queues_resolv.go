package api

type ResolverPersonalQueues struct {
	s TypePersonalQueues
}

func (R *ResolverPersonalQueues) Set(s TypePersonalQueues) {
	R.s = s
}

func (R ResolverPersonalQueues) UserId() int32 {
	return R.s.UserId
}

func (R ResolverPersonalQueues) QueueId() *int32 {
	return R.s.QueueId
}

func GenGqlTypePersonalQueues(extra string) string {
	return "type PersonalQueues { " + extra + `
	UserId: Int,
	QueueId: Int,
	}`
}
