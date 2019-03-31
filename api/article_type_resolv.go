package api

type ResolverArticleType struct {
	s TypeArticleType
}

func (R *ResolverArticleType) Set(s TypeArticleType) {
	R.s = s
}

func (R ResolverArticleType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverArticleType) Name() *string {
	return R.s.Name
}

func (R ResolverArticleType) Comments() *string {
	return R.s.Comments
}

func (R ResolverArticleType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverArticleType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverArticleType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverArticleType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverArticleType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeArticleType(extra string) string {
	return "type ArticleType { " + extra + `
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
