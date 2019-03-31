package api

type ResolverGroups struct {
	s TypeGroups
}

func (R *ResolverGroups) Set(s TypeGroups) {
	R.s = s
}

func (R ResolverGroups) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverGroups) Name() *string {
	return R.s.Name
}

func (R ResolverGroups) Comments() *string {
	return R.s.Comments
}

func (R ResolverGroups) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverGroups) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverGroups) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverGroups) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverGroups) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeGroups(extra string) string {
	return "type Groups { " + extra + `
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
