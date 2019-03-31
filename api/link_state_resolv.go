package api

type ResolverLinkState struct {
	s TypeLinkState
}

func (R *ResolverLinkState) Set(s TypeLinkState) {
	R.s = s
}

func (R ResolverLinkState) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverLinkState) Name() *string {
	return R.s.Name
}

func (R ResolverLinkState) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverLinkState) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverLinkState) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverLinkState) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverLinkState) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeLinkState(extra string) string {
	return "type LinkState { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
