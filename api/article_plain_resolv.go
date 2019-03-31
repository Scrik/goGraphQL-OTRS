package api

type ResolverArticlePlain struct {
	s TypeArticlePlain
}

func (R *ResolverArticlePlain) Set(s TypeArticlePlain) {
	R.s = s
}

func (R ResolverArticlePlain) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverArticlePlain) ArticleId() *string {
	str := fmt.Sprintf("%d", R.s.ArticleId)
	return &str
}

func (R *ResolverArticlePlain) Body() *string {
	str := string(*R.s.Body)
	return &str
}

func (R *ResolverArticlePlain) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverArticlePlain) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverArticlePlain) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverArticlePlain) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeArticlePlain(extra string) string {
	return "type ArticlePlain { " + extra + `
	Id: String,
	ArticleId: String,
	Body: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
