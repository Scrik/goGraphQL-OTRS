package api

type ResolverGroupCustomerUser struct {
	s TypeGroupCustomerUser
}

func (R *ResolverGroupCustomerUser) Set(s TypeGroupCustomerUser) {
	R.s = s
}

func (R ResolverGroupCustomerUser) UserId() string {
	return R.s.UserId
}

func (R ResolverGroupCustomerUser) GroupId() *int32 {
	return R.s.GroupId
}

func (R ResolverGroupCustomerUser) PermissionKey() *string {
	return R.s.PermissionKey
}

func (R ResolverGroupCustomerUser) PermissionValue() *int32 {
	i := int32(*R.s.PermissionValue)
	return &i
}

func (R *ResolverGroupCustomerUser) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverGroupCustomerUser) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverGroupCustomerUser) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverGroupCustomerUser) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeGroupCustomerUser(extra string) string {
	return "type GroupCustomerUser { " + extra + `
	UserId: String,
	GroupId: Int,
	PermissionKey: String,
	PermissionValue: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
