package api

type ResolverServiceCustomerUser struct {
	s TypeServiceCustomerUser
}

func (R *ResolverServiceCustomerUser) Set(s TypeServiceCustomerUser) {
	R.s = s
}

func (R ResolverServiceCustomerUser) CustomerUserLogin() string {
	return R.s.CustomerUserLogin
}

func (R ResolverServiceCustomerUser) ServiceId() *int32 {
	return R.s.ServiceId
}

func (R *ResolverServiceCustomerUser) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverServiceCustomerUser) CreateBy() *int32 {
	return R.s.CreateBy
}

func GenGqlTypeServiceCustomerUser(extra string) string {
	return "type ServiceCustomerUser { " + extra + `
	CustomerUserLogin: String,
	ServiceId: Int,
	CreateTime: String,
	CreateBy: Int,
	}`
}
