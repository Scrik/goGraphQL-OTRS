package api

type ResolverAutoResponseType struct {
	s TypeAutoResponseType
}

func (R *ResolverAutoResponseType) Set(s TypeAutoResponseType) {
	R.s = s
}

func (R ResolverAutoResponseType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverAutoResponseType) Name() *string {
	return R.s.Name
}

func (R ResolverAutoResponseType) Comments() *string {
	return R.s.Comments
}

func (R ResolverAutoResponseType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverAutoResponseType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverAutoResponseType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverAutoResponseType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverAutoResponseType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeAutoResponseType(extra string) string {
	return "type AutoResponseType { " + extra + `
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
