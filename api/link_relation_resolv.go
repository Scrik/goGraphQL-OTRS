package api

type ResolverLinkRelation struct {
	s TypeLinkRelation
}

func (R *ResolverLinkRelation) Set(s TypeLinkRelation) {
	R.s = s
}

func (R ResolverLinkRelation) SourceObjectId() int8 {
	return R.s.SourceObjectId
}

func (R ResolverLinkRelation) SourceKey() *string {
	return R.s.SourceKey
}

func (R ResolverLinkRelation) TargetObjectId() *int32 {
	i := int32(*R.s.TargetObjectId)
	return &i
}

func (R ResolverLinkRelation) TargetKey() *string {
	return R.s.TargetKey
}

func (R ResolverLinkRelation) TypeId() *int32 {
	i := int32(*R.s.TypeId)
	return &i
}

func (R ResolverLinkRelation) StateId() *int32 {
	i := int32(*R.s.StateId)
	return &i
}

func (R *ResolverLinkRelation) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverLinkRelation) CreateBy() *int32 {
	return R.s.CreateBy
}

func GenGqlTypeLinkRelation(extra string) string {
	return "type LinkRelation { " + extra + `
	SourceObjectId: Int,
	SourceKey: String,
	TargetObjectId: Int,
	TargetKey: String,
	TypeId: Int,
	StateId: Int,
	CreateTime: String,
	CreateBy: Int,
	}`
}
