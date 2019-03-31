package api

type ResolverArticleSenderType struct {
	s TypeArticleSenderType
}

func (R *ResolverArticleSenderType) Set(s TypeArticleSenderType) {
	R.s = s
}

func (R ResolverArticleSenderType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverArticleSenderType) Name() *string {
	return R.s.Name
}

func (R ResolverArticleSenderType) Comments() *string {
	return R.s.Comments
}

func (R ResolverArticleSenderType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverArticleSenderType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverArticleSenderType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverArticleSenderType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverArticleSenderType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeArticleSenderType(extra string) string {
	return "type ArticleSenderType { " + extra + `
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
