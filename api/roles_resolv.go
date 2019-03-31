package api

type ResolverRoles struct {
	s TypeRoles
}

func (R *ResolverRoles) Set(s TypeRoles) {
	R.s = s
}

func (R ResolverRoles) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverRoles) Name() *string {
	return R.s.Name
}

func (R ResolverRoles) Comments() *string {
	return R.s.Comments
}

func (R ResolverRoles) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverRoles) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverRoles) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverRoles) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverRoles) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeRoles(extra string) string {
	return "type Roles { " + extra + `
	Id: Int,
	Name: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
