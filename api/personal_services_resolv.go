package api

type ResolverPersonalServices struct {
	s TypePersonalServices
}

func (R *ResolverPersonalServices) Set(s TypePersonalServices) {
	R.s = s
}

func (R ResolverPersonalServices) UserId() int32 {
	return R.s.UserId
}

func (R ResolverPersonalServices) ServiceId() *int32 {
	return R.s.ServiceId
}

func GenGqlTypePersonalServices(extra string) string {
	return "type PersonalServices { " + extra + `
	UserId: Int,
	ServiceId: Int,
	}`
}
