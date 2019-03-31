package api

type ResolverRoleUser struct {
	s TypeRoleUser
}

func (R *ResolverRoleUser) Set(s TypeRoleUser) {
	R.s = s
}

func (R ResolverRoleUser) UserId() int32 {
	return R.s.UserId
}

func (R ResolverRoleUser) RoleId() *int32 {
	return R.s.RoleId
}

func (R *ResolverRoleUser) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverRoleUser) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverRoleUser) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverRoleUser) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeRoleUser(extra string) string {
	return "type RoleUser { " + extra + `
	UserId: Int,
	RoleId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
