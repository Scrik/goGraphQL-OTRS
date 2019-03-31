package api

type ResolverLinkType struct {
	s TypeLinkType
}

func (R *ResolverLinkType) Set(s TypeLinkType) {
	R.s = s
}

func (R ResolverLinkType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverLinkType) Name() *string {
	return R.s.Name
}

func (R ResolverLinkType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverLinkType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverLinkType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverLinkType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverLinkType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeLinkType(extra string) string {
	return "type LinkType { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
