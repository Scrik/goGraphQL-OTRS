package api

type ResolverGroupUser struct {
	s TypeGroupUser
}

func (R *ResolverGroupUser) Set(s TypeGroupUser) {
	R.s = s
}

func (R ResolverGroupUser) UserId() int32 {
	return R.s.UserId
}

func (R ResolverGroupUser) GroupId() *int32 {
	return R.s.GroupId
}

func (R ResolverGroupUser) PermissionKey() *string {
	return R.s.PermissionKey
}

func (R ResolverGroupUser) PermissionValue() *int32 {
	i := int32(*R.s.PermissionValue)
	return &i
}

func (R *ResolverGroupUser) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverGroupUser) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverGroupUser) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverGroupUser) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeGroupUser(extra string) string {
	return "type GroupUser { " + extra + `
	UserId: Int,
	GroupId: Int,
	PermissionKey: String,
	PermissionValue: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
