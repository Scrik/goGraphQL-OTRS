package api

type ResolverGroupRole struct {
	s TypeGroupRole
}

func (R *ResolverGroupRole) Set(s TypeGroupRole) {
	R.s = s
}

func (R ResolverGroupRole) RoleId() int32 {
	return R.s.RoleId
}

func (R ResolverGroupRole) GroupId() *int32 {
	return R.s.GroupId
}

func (R ResolverGroupRole) PermissionKey() *string {
	return R.s.PermissionKey
}

func (R ResolverGroupRole) PermissionValue() *int32 {
	i := int32(*R.s.PermissionValue)
	return &i
}

func (R *ResolverGroupRole) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverGroupRole) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverGroupRole) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverGroupRole) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeGroupRole(extra string) string {
	return "type GroupRole { " + extra + `
	RoleId: Int,
	GroupId: Int,
	PermissionKey: String,
	PermissionValue: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
